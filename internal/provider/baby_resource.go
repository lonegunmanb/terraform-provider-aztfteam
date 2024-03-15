// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	myvalidator "github.com/lonegunmanb/terraform-provider-aztfteam/internal/validator"
	"github.com/magodo/age"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &BabyResource{}
	_ resource.ResourceWithImportState = &BabyResource{}
)

func NewBabyResource() resource.Resource {
	return &BabyResource{}
}

// BabyResource defines the resource implementation.
type BabyResource struct {
	client *http.Client
}

// BabyResourceModel describes the resource data model.
type BabyResourceModel struct {
	Age              types.Int64  `tfsdk:"age"`
	Agility          types.Number `tfsdk:"agility"`
	Birthday         types.String `tfsdk:"birthday"`
	Charisma         types.Number `tfsdk:"charisma"`
	Endurance        types.Number `tfsdk:"endurance"`
	Id               types.String `tfsdk:"id"`
	Intelligence     types.Int64  `tfsdk:"intelligence"`
	Luck             types.Number `tfsdk:"luck"`
	Name             types.String `tfsdk:"name"`
	Strength         types.Number `tfsdk:"strength"`
	Perception       types.Number `tfsdk:"perception"`
	Tags             types.Map    `tfsdk:"tags"`
	BiologicalGender types.String `tfsdk:"biological_gender"`
}

func (r *BabyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_baby"
}

func (r *BabyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Baby resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Internal identifier",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Baby's name",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"birthday": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The birthday in RFC3339 format",
				Validators: []validator.String{
					myvalidator.StringIsParsable("birthday", func(s string) error {
						_, err := time.Parse(time.RFC3339, s)
						return err
					}),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"age": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Baby's age",
			},
			"strength": schema.NumberAttribute{
				Computed:            true,
				MarkdownDescription: "Baby's strength",
			},
			"endurance": schema.NumberAttribute{
				Computed:            true,
				MarkdownDescription: "Baby's endurance",
			},
			"agility": schema.NumberAttribute{
				Computed:            true,
				MarkdownDescription: "Baby's agility",
			},
			"luck": schema.NumberAttribute{
				Computed:            true,
				MarkdownDescription: "Baby's luck",
			},
			"perception": schema.NumberAttribute{
				Computed:            true,
				MarkdownDescription: "Baby's perception",
			},
			"charisma": schema.NumberAttribute{
				Computed:            true,
				MarkdownDescription: "Baby's charisma",
			},
			"intelligence": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Baby's intelligence",
			},
			"biological_gender": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Baby's biological gender",
			},
			"tags": schema.MapAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: "Baby's tags",
			},
		},
	}
}

func (r *BabyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*http.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *BabyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BabyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create example, got error: %s", err))
	//     return
	// }

	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.
	data.Id = types.StringValue(uuid.NewString())

	if data.Birthday.IsUnknown() {
		data.Birthday = types.StringValue(time.Now().Format(time.RFC3339))
	}
	birthDay, err := time.Parse(time.RFC3339, data.Birthday.ValueString())
	if err != nil {
		// This shouldn't happen as the schema has validated it, while we still test against this for completeness.
		panic(fmt.Errorf("failed to parse birthday %q: %v", data.Birthday, err))
	}
	data.Age = types.Int64Value(int64(age.Age(birthDay, time.Now())))

	// According to https://help.bethesda.net/#en/answer/44321: "Each individual attribute has the potential to reach a maximum total of 15 points assigned."
	data.Agility = types.NumberValue(big.NewFloat(float64(10 + rand.Int31n(6))))
	data.Endurance = types.NumberValue(big.NewFloat(float64(10 + rand.Int31n(6))))
	data.Intelligence = types.Int64Value(int64(100 + rand.Int31n(40)))
	data.Luck = types.NumberValue(big.NewFloat(float64(10 + rand.Int31n(6))))
	data.Charisma = types.NumberValue(big.NewFloat(float64(10 + rand.Int31n(6))))
	data.Strength = types.NumberValue(big.NewFloat(float64(10 + rand.Int31n(6))))
	data.Perception = types.NumberValue(big.NewFloat(float64(10 + rand.Int31n(6))))
	data.Tags, _ = types.MapValue(types.StringType, map[string]attr.Value{
		"blessed_by": types.StringValue("terraform engineering China team"),
	})

	genders := []string{"boy", "girl"}
	data.BiologicalGender = types.StringValue(genders[rand.Int()%len(genders)])

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a baby")

	time.Sleep(100 * time.Second)
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BabyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BabyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	// As long as the baby is created naturally, the birthday shall always be available.
	if !data.Birthday.IsNull() {
		birthDay, err := time.Parse(time.RFC3339, data.Birthday.ValueString())
		if err != nil {
			// This shouldn't happen as the schema has validated it, while we still test against this for completeness.
			panic(fmt.Errorf("failed to parse birthday %q: %v", data.Birthday, err))
		}
		data.Age = types.Int64Value(int64(age.Age(birthDay, time.Now())))
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BabyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data BabyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BabyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Error(ctx, fmt.Sprintf("Baby can not be deleted!"))
}

func (r *BabyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

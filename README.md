# Terraform Provider: Baby Resource with SPECIAL System

The Terraform provider for the baby resource is a fun and engaging way to simulate the birth of a baby. In the future, we plan to enhance this resource by incorporating the SPECIAL system from the Fallout game series.

## SPECIAL System

The SPECIAL system is a character development system used in the Fallout games. The name "SPECIAL" is an acronym standing for the primary attributes in the system: Strength, Perception, Endurance, Charisma, Intelligence, Agility, and Luck. Each of these attributes contributes to the skills and abilities of player characters within the game.

## Baby Resource with SPECIAL System

In the context of the baby resource, the SPECIAL system will be used to assign attributes to the baby at the time of its creation. Each baby will be assigned a value for each of the seven SPECIAL attributes. These attributes will be randomly generated within a certain range, similar to the game mechanics in Fallout.

Here's what the baby resource would look like with the SPECIAL system:

```hcl
resource "aztfteam_baby" "one" {
  name = "neo"

  strength     = 10
  perception   = 12
  endurance    = 11
  charisma     = 14
  intelligence = 13
  agility      = 15
  luck         = 10
}
```

In this example, the baby named "example" has been assigned the following SPECIAL attributes: Strength of 10, Perception of 12, Endurance of 11, Charisma of 14, Intelligence of 13, Agility of 15, and Luck of 10.

These attributes will be computed and stored in the Terraform state. They can be referenced in other resources or outputs, providing a fun and interactive way to simulate the birth and growth of a baby in your Terraform configuration.

Please note that these attributes are for simulation purposes only and do not have any impact on the actual operation of the Terraform provider.
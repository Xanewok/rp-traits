// Roll order for the fighters
f["class"] = pickTrait(f, fighterClasses, src)
f["body"] = pickTrait(f, fighterBodies, src)
switch f["class"].name {
case "Slayer":
    f["weapon"] = pickTrait(f, fighterSwords, src)
case "Paladin":
    f["weapon"] = pickTrait(f, fighterHammers, src)
case "Lancer":
    f["weapon"] = pickTrait(f, fighterLances, src)
case "Archer":
    f["weapon"] = pickTrait(f, fighterBows, src)
case "Mage":
    f["weapon"] = pickTrait(f, fighterWands, src)
case "Assassin":
    f["weapon"] = pickTrait(f, fighterThrown, src)
}
f["hair"] = pickTrait(f, fighterHairs, src)
f["hairColor"] = pickTrait(f, fighterHairColors, src)
f["back"] = pickTrait(f, fighterBacks, src)
f["aura"] = pickTrait(f, fighterAuras, src)
f["top"] = pickTrait(f, fighterTops, src)
f["bottom"] = pickTrait(f, fighterBottoms, src)

// Clash/force logic
func pickTrait(fighter Fighter, traits []Trait, src *rand.Rand) Trait {

	// For all traits in hero, if the hero trait feature has a forced trait,
	// and the forced trait matches the given traitset, force the trait.

	for _, trait := range fighter {
		if len(trait.force) != 0 {
			if trait.force[0].feature == traits[0].feature {
				return trait.force[0]
			}
		}
	}

	// Make array of current clash traits, and declare an array of filtered traits.
	clashes := buildFighterClashes(fighter)
	filteredTraits := traits

	// For each clash, and for each given trait in the trait set,
	// if the clash feature matches the feature of the trait,
	// and their id's are the same, remove the clash trait from
	// the set of filtered traits.
	if len(fighter) > 0 {
		for _, c := range clashes {
			for _, t := range traits {
				if c.feature == traits[0].feature {
					if c.id == t.id {
						filteredTraits = remove(filteredTraits, c.id)
					}
				}
			}
		}
	}

	// Make array of weighted choices
	var choices []wr.Choice
	for _, s := range filteredTraits {
		choices = append(choices, wr.Choice{Item: s.id, Weight: uint(s.weight)})
	}
	chooser, _ := wr.NewChooser(choices...)
	result := chooser.PickSource(src).(int)

	for _, t := range filteredTraits {
		if result == t.id {
			return t
		}
	}

	return Trait{id: 0, name: "NONE"}
}

// Actual values
type Trait struct {
	feature    string
	id         int
	name       string
	weight     int
	properties []string
	clashes    []Trait
	force      []Trait
}

var fighterBows = []Trait{
	{
		feature:    "weapon",
		id:         0,
		name:       "Wooden",
		weight:     22,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         1,
		name:       "Longbow",
		weight:     15,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         2,
		name:       "Compound",
		weight:     13,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         3,
		name:       "Quartz Shot",
		weight:     12,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         4,
		name:       "Waterweaver",
		weight:     10,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         5,
		name:       "Faerie",
		weight:     9,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         6,
		name:       "Black Hole",
		weight:     7,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         7,
		name:       "Beelzebub",
		weight:     6,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         8,
		name:       "Raphael",
		weight:     5,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         9,
		name:       "Gandiva",
		weight:     1,
		properties: []string{"hasBack"},
	},
}

var fighterWands = []Trait{
	{
		feature:    "weapon",
		id:         0,
		name:       "Elder Staff",
		weight:     22,
		properties: []string{"hasBack", "mage"},
	},
	{
		feature:    "weapon",
		id:         1,
		name:       "Dream Catcher",
		weight:     15,
		properties: []string{"hasBack", "mage"},
	},
	{
		feature:    "weapon",
		id:         2,
		name:       "Antenna",
		weight:     13,
		properties: []string{"hasBack", "mage"},
	},
	{
		feature:    "weapon",
		id:         3,
		name:       "Chrysoberyl",
		weight:     12,
		properties: []string{"hasBack", "mage"},
	},
	{
		feature:    "weapon",
		id:         4,
		name:       "Monsoon",
		weight:     10,
		properties: []string{"hasBack", "mage"},
	},
	{
		feature:    "weapon",
		id:         5,
		name:       "Sands of Time",
		weight:     9,
		properties: []string{"hasBack", "mage"},
	},
	{
		feature:    "weapon",
		id:         6,
		name:       "Eclipse",
		weight:     7,
		properties: []string{"hasBack", "mage"},
	},
	{
		feature:    "weapon",
		id:         7,
		name:       "Original Sin",
		weight:     6,
		properties: []string{"hasBack", "mage"},
	},
	{
		feature:    "weapon",
		id:         8,
		name:       "Gabriel",
		weight:     5,
		properties: []string{"hasBack", "mage"},
	},
	{
		feature:    "weapon",
		id:         9,
		name:       "Yata",
		weight:     1,
		properties: []string{"hasBack", "mage"},
	},
}

var fighterThrown = []Trait{
	{
		feature:    "weapon",
		id:         0,
		name:       "Knife",
		weight:     22,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         1,
		name:       "Shuriken",
		weight:     15,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         2,
		name:       "Syringe",
		weight:     13,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         3,
		name:       "Cystal Fragments",
		weight:     12,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         4,
		name:       "Cinders",
		weight:     10,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         5,
		name:       "Sacrificial",
		weight:     9,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         6,
		name:       "Stars",
		weight:     7,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         7,
		name:       "Deception",
		weight:     6,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         8,
		name:       "Heaven",
		weight:     5,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         9,
		name:       "Parazonium",
		weight:     1,
		properties: []string{"hasBack"},
	},
}

var fighterSwords = []Trait{
	{
		feature:    "weapon",
		id:         0,
		name:       "Iron Sword",
		weight:     22,
		properties: []string{""},
	},
	{
		feature:    "weapon",
		id:         1,
		name:       "Spiked Club",
		weight:     15,
		properties: []string{""},
	},
	{
		feature:    "weapon",
		id:         2,
		name:       "Boxcutter",
		weight:     13,
		properties: []string{""},
	},
	{
		feature:    "weapon",
		id:         3,
		name:       "Jadecutter",
		weight:     12,
		properties: []string{""},
	},
	{
		feature:    "weapon",
		id:         4,
		name:       "Caustic",
		weight:     10,
		properties: []string{""},
	},
	{
		feature:    "weapon",
		id:         5,
		name:       "Wandering Blade",
		weight:     9,
		properties: []string{""},
	},
	{
		feature:    "weapon",
		id:         6,
		name:       "Lunar",
		weight:     7,
		properties: []string{""},
	},
	{
		feature:    "weapon",
		id:         7,
		name:       "Lucifer",
		weight:     6,
		properties: []string{""},
	},
	{
		feature:    "weapon",
		id:         8,
		name:       "Uriel",
		weight:     5,
		properties: []string{""},
	},
	{
		feature:    "weapon",
		id:         9,
		name:       "Joyeuse",
		weight:     1,
		properties: []string{""},
	},
}

var fighterLances = []Trait{
	{
		feature:    "weapon",
		id:         0,
		name:       "Javelin",
		weight:     22,
		properties: []string{""},
	},
	{
		feature:    "weapon",
		id:         1,
		name:       "Lance",
		weight:     15,
		properties: []string{""},
	},
	{
		feature:    "weapon",
		id:         2,
		name:       "Gunlance",
		weight:     13,
		properties: []string{""},
	},
	{
		feature:    "weapon",
		id:         3,
		name:       "Skypiercer",
		weight:     12,
		properties: []string{""},
	},
	{
		feature:    "weapon",
		id:         4,
		name:       "Glacial Lance",
		weight:     10,
		properties: []string{""},
	},
	{
		feature:    "weapon",
		id:         5,
		name:       "Banner",
		weight:     9,
		properties: []string{""},
	},
	{
		feature:    "weapon",
		id:         6,
		name:       "Scorpius",
		weight:     7,
		properties: []string{""},
	},
	{
		feature:    "weapon",
		id:         7,
		name:       "Harbinger",
		weight:     6,
		properties: []string{""},
	},
	{
		feature:    "weapon",
		id:         8,
		name:       "Michael",
		weight:     5,
		properties: []string{""},
	},
	{
		feature:    "weapon",
		id:         9,
		name:       "Gae Bolg",
		weight:     1,
		properties: []string{""},
	},
}

var fighterHammers = []Trait{
	{
		feature:    "weapon",
		id:         0,
		name:       "Mace",
		weight:     22,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         1,
		name:       "Smith",
		weight:     15,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         2,
		name:       "Chainsaw",
		weight:     13,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         3,
		name:       "Onyx",
		weight:     12,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         4,
		name:       "Monolith",
		weight:     10,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         5,
		name:       "Scythe",
		weight:     9,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         6,
		name:       "Jupiter",
		weight:     7,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         7,
		name:       "Leviathan",
		weight:     6,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         8,
		name:       "True Cross",
		weight:     5,
		properties: []string{"hasBack"},
	},
	{
		feature:    "weapon",
		id:         9,
		name:       "Mjolnir",
		weight:     1,
		properties: []string{"hasBack"},
	},
}

var fighterBacks = []Trait{
	{
		feature:    "back",
		id:         0,
		name:       "None",
		weight:     80,
		properties: []string{""},
	},
	{
		feature:    "back",
		id:         1,
		name:       "Ribbon",
		weight:     10,
		properties: []string{""},
	},
	{
		feature:    "back",
		id:         2,
		name:       "Fairy",
		weight:     5,
		properties: []string{""},
	},
	{
		feature:    "back",
		id:         3,
		name:       "Best Friend",
		weight:     5,
		properties: []string{""},
	},
	{
		feature:    "back",
		id:         4,
		name:       "Windup",
		weight:     10,
		properties: []string{""},
	},
	{
		feature:    "back",
		id:         5,
		name:       "Yellow Flag",
		weight:     10,
		properties: []string{""},
	},
	{
		feature:    "back",
		id:         6,
		name:       "Red Flag",
		weight:     10,
		properties: []string{""},
	},
	{
		feature:    "back",
		id:         7,
		name:       "Danger Coat",
		weight:     2,
		properties: []string{""},
	},
	{
		feature:    "back",
		id:         8,
		name:       "Syringe",
		weight:     10,
		properties: []string{""},
	},
	{
		feature:    "back",
		id:         9,
		name:       "Dragon Tail",
		weight:     10,
		properties: []string{""},
	},
	{
		feature:    "back",
		id:         10,
		name:       "Reaper Cape",
		weight:     1,
		properties: []string{""},
	},
	{
		feature:    "back",
		id:         11,
		name:       "Stuntman",
		weight:     25,
		properties: []string{""},
	},
	{
		feature:    "back",
		id:         12,
		name:       "Blue Balloon",
		weight:     25,
		properties: []string{""},
	},
	{
		feature:    "back",
		id:         13,
		name:       "Pink Balloon",
		weight:     25,
		properties: []string{""},
	},
	{
		feature:    "back",
		id:         14,
		name:       "Yellow Balloon",
		weight:     25,
		properties: []string{""},
	},
	{
		feature:    "back",
		id:         15,
		name:       "Clock",
		weight:     2,
		properties: []string{""},
	},
	{
		feature:    "back",
		id:         16,
		name:       "Black Tassels",
		weight:     5,
		properties: []string{""},
	},
	{
		feature:    "back",
		id:         17,
		name:       "White Cape",
		weight:     10,
		properties: []string{""},
	},
	{
		feature:    "back",
		id:         18,
		name:       "Green Cape",
		weight:     10,
		properties: []string{""},
	},
	{
		feature:    "back",
		id:         19,
		name:       "Winter Cloak",
		weight:     10,
		properties: []string{""},
	},
}
var fighterHairs = []Trait{
	{
		feature:    "hair",
		id:         0,
		name:       "Overgrown",
		weight:     500,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         1,
		name:       "Ponytail",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         2,
		name:       "Twinkle",
		weight:     63,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         3,
		name:       "Vocaloid",
		weight:     63,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         4,
		name:       "Soprano",
		weight:     150,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         5,
		name:       "Mid Part Updo",
		weight:     500,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         6,
		name:       "Spikehawk",
		weight:     500,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         7,
		name:       "Preppy",
		weight:     500,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         8,
		name:       "Bed Head",
		weight:     500,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         9,
		name:       "Overgrown",
		weight:     500,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         10,
		name:       "Bow Hair",
		weight:     150,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         11,
		name:       "Antenna",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         12,
		name:       "Emo Midcut",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         13,
		name:       "Samurai Cut",
		weight:     150,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         14,
		name:       "Twintails",
		weight:     500,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         15,
		name:       "Commander",
		weight:     500,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         16,
		name:       "Updo",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         17,
		name:       "Rockstar",
		weight:     10,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         18,
		name:       "Buzzcut",
		weight:     500,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         19,
		name:       "Villain Hair",
		weight:     63,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         20,
		name:       "Female Buns",
		weight:     500,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         21,
		name:       "Short Bob",
		weight:     500,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         22,
		name:       "Short Buzz Cut",
		weight:     500,
		properties: []string{""},
	},
}
var longHairs = []Trait{
	{
		feature:    "hair",
		id:         0,
		name:       "Overgrown",
		weight:     10,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         1,
		name:       "Ponytail",
		weight:     10,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         2,
		name:       "Twinkle",
		weight:     10,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         3,
		name:       "Vocaloid",
		weight:     10,
		properties: []string{""},
	},
	{
		feature:    "hair",
		id:         4,
		name:       "Soprano",
		weight:     10,
		properties: []string{""},
	},
}
var fighterHairColors = []Trait{
	{
		feature: "hairColors",
		id:      0,
		name:    "Blue",
		weight:  150,
	},
	{
		feature: "hairColors",
		id:      1,
		name:    "Pink",
		weight:  50,
	},
	{
		feature: "hairColors",
		id:      2,
		name:    "Green",
		weight:  150,
	},
	{
		feature: "hairColors",
		id:      3,
		name:    "Violet",
		weight:  125,
	},
	{
		feature: "hairColors",
		id:      4,
		name:    "Red",
		weight:  50,
	},
	{
		feature: "hairColors",
		id:      5,
		name:    "Black",
		weight:  150,
	},
	{
		feature: "hairColors",
		id:      6,
		name:    "White",
		weight:  25,
	},
	{
		feature: "hairColors",
		id:      7,
		name:    "Blonde",
		weight:  150,
	},
	{
		feature: "hairColors",
		id:      8,
		name:    "Brown",
		weight:  150,
	},
}

var fighterAuras = []Trait{

	{
		feature:    "aura",
		id:         0,
		name:       "None",
		weight:     8750,
		properties: []string{""},
	},
	{
		feature:    "aura",
		id:         1,
		name:       "IV",
		weight:     100,
		properties: []string{"hasFront"},
	},
	{
		feature:    "aura",
		id:         2,
		name:       "Puppet",
		weight:     50,
		properties: []string{"hasFront"},
	},
	{
		feature:    "aura",
		id:         3,
		name:       "Sword",
		weight:     100,
		properties: []string{""},
	},
	{
		feature:    "aura",
		id:         4,
		name:       "Sun",
		weight:     75,
		properties: []string{""},
	},
	{
		feature:    "aura",
		id:         5,
		name:       "Feathers",
		weight:     100,
		properties: []string{""},
	},
	{
		feature:    "aura",
		id:         6,
		name:       "Flowers",
		weight:     75,
		properties: []string{""},
	},
	{
		feature:    "aura",
		id:         7,
		name:       "Chess",
		weight:     100,
		properties: []string{""},
	},
	{
		feature:    "aura",
		id:         8,
		name:       "Snowflakes",
		weight:     75,
		properties: []string{""},
	},
	{
		feature:    "aura",
		id:         9,
		name:       "Gears",
		weight:     100,
		properties: []string{""},
	},
	{
		feature:    "aura",
		id:         10,
		name:       "Tree",
		weight:     75,
		properties: []string{""},
	},
	{
		feature:    "aura",
		id:         11,
		name:       "Chains",
		weight:     50,
		properties: []string{""},
	},
	{
		feature:    "aura",
		id:         12,
		name:       "Spotlight",
		weight:     75,
		properties: []string{""},
	},
	{
		feature:    "aura",
		id:         13,
		name:       "Feathers 2",
		weight:     100,
		properties: []string{""},
	},
	{
		feature:    "aura",
		id:         14,
		name:       "Heart",
		weight:     100,
		properties: []string{""},
	},
	{
		feature:    "aura",
		id:         15,
		name:       "Moon",
		weight:     75,
		properties: []string{""},
	},
}

var fighterClasses = []Trait{
	{
		feature: "Class",
		id:      0,
		name:    "Slayer",
		weight:  10,
		clashes: []Trait{
			{
				feature:    "back",
				id:         18,
				name:       "Green Cape",
				weight:     10,
				properties: []string{""},
			},
			{
				feature:    "back",
				id:         16,
				name:       "Black Tassels",
				weight:     10,
				properties: []string{""},
			},
			{
				feature:    "back",
				id:         7,
				name:       "Danger Coat",
				weight:     10,
				properties: []string{""},
			},
		},
	},
	{
		feature: "Class",
		id:      1,
		name:    "Paladin",
		weight:  10,
	},
	{
		feature: "Class",
		id:      2,
		name:    "Lancer",
		weight:  10,
	},
	{
		feature: "Class",
		id:      3,
		name:    "Archer",
		weight:  10,
		clashes: []Trait{
			{
				feature:    "back",
				id:         18,
				name:       "Green Cape",
				weight:     10,
				properties: []string{""},
			},
			{
				feature:    "back",
				id:         16,
				name:       "Black Tassels",
				weight:     10,
				properties: []string{""},
			},
			{
				feature:    "back",
				id:         7,
				name:       "Danger Coat",
				weight:     10,
				properties: []string{""},
			},
		},
	},
	{
		feature: "Class",
		id:      4,
		name:    "Mage",
		weight:  10,
	},
	{
		feature: "class",
		id:      5,
		name:    "Assassin",
		weight:  10,
	},
}

var fighterBottoms = []Trait{
	{
		feature:    "bottom",
		id:         0,
		name:       "Black Slacks",
		weight:     9,
		properties: []string{""},
	},
	{
		feature:    "bottom",
		id:         1,
		name:       "Ragged Pants",
		weight:     9,
		properties: []string{""},
	},
	{
		feature:    "bottom",
		id:         2,
		name:       "Blue Skirt",
		weight:     9,
		properties: []string{""},
	},
	{
		feature:    "bottom",
		id:         3,
		name:       "White Pants",
		weight:     9,
		properties: []string{""},
	},
	{
		feature:    "bottom",
		id:         4,
		name:       "White Shorts",
		weight:     9,
		properties: []string{""},
	},
	{
		feature:    "bottom",
		id:         5,
		name:       "Blue Training Shorts",
		weight:     9,
		properties: []string{""},
	},
	{
		feature:    "bottom",
		id:         6,
		name:       "Red Training Shorts",
		weight:     9,
		properties: []string{""},
	},
	{
		feature:    "bottom",
		id:         7,
		name:       "Black Edgy Pants",
		weight:     5,
		properties: []string{""},
	},
	{
		feature:    "bottom",
		id:         8,
		name:       "Disco Pants",
		weight:     5,
		properties: []string{""},
	},
	{
		feature:    "bottom",
		id:         9,
		name:       "Gold Sparkle Pleats",
		weight:     5,
		properties: []string{""},
	},
	{
		feature:    "bottom",
		id:         10,
		name:       "White Skirt",
		weight:     5,
		properties: []string{""},
	},
	{
		feature:    "bottom",
		id:         11,
		name:       "Armored Pants",
		weight:     5,
		properties: []string{""},
	},
	{
		feature:    "bottom",
		id:         12,
		name:       "Black Sneak Pants",
		weight:     3,
		properties: []string{""},
	},
	{
		feature:    "bottom",
		id:         13,
		name:       "Tan Sneak Pants",
		weight:     3,
		properties: []string{""},
	},
	{
		feature:    "bottom",
		id:         14,
		name:       "Martial Arts Pants",
		weight:     3,
		properties: []string{""},
	},
	{
		feature:    "bottom",
		id:         15,
		name:       "Boxers",
		weight:     3,
		properties: []string{""},
	},
	{
		feature:    "bottom",
		id:         16,
		name:       "Farmer Apron",
		weight:     0,
		properties: []string{""},
	},
	{
		feature:    "bottom",
		id:         17,
		name:       "Thrifty Getup",
		weight:     0,
		properties: []string{""},
	},
	{
		feature:    "bottom",
		id:         18,
		name:       "None",
		weight:     0,
		properties: []string{""},
	},
}

var fighterTops = []Trait{
	{
		feature:    "top",
		id:         0,
		name:       "Heart Tank",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         1,
		name:       "Track Jacket",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         2,
		name:       "Pocket Tee",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         3,
		name:       "Pink Showstopper",
		weight:     300,
		properties: []string{""},
		force: []Trait{
			{
				feature:    "bottom",
				id:         18,
				name:       "None",
				weight:     0,
				properties: []string{""},
			},
		},
	},
	{
		feature:    "top",
		id:         4,
		name:       "Ab T-Shirt",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         5,
		name:       "Bunny Jersey",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         6,
		name:       "Lover Sweater",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         7,
		name:       "Red Checker",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         8,
		name:       "Green Archer’s Hood",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         9,
		name:       "Red Archer’s Hood",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         10,
		name:       "Lei",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         11,
		name:       "Red Vest",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         12,
		name:       "Black Vest",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         13,
		name:       "Baseball Shirt",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         14,
		name:       "Parka",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         15,
		name:       "Life Jacket",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         16,
		name:       "Tank Top",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         17,
		name:       "Sailor Outfit",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         18,
		name:       "Casual Suit",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         19,
		name:       "Star Sweater",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         20,
		name:       "Heart Tee",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         21,
		name:       "Bunny T-Shirt",
		weight:     300,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         22,
		name:       "Abstract Hood",
		weight:     150,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         23,
		name:       "Gunslinger",
		weight:     150,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         24,
		name:       "Olympian’s Robe",
		weight:     150,
		properties: []string{""},
		force: []Trait{
			{
				feature:    "bottom",
				id:         18,
				name:       "None",
				weight:     0,
				properties: []string{""},
			},
		},
	},
	{
		feature:    "top",
		id:         25,
		name:       "Dojo Casual",
		weight:     150,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         26,
		name:       "Desert Traveler",
		weight:     150,
		properties: []string{""},
		force: []Trait{
			{
				feature:    "bottom",
				id:         18,
				name:       "None",
				weight:     0,
				properties: []string{""},
			},
		},
	},
	{
		feature:    "top",
		id:         27,
		name:       "Soccer Jersey",
		weight:     150,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         28,
		name:       "Distinguished Coat",
		weight:     150,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         29,
		name:       "Lab Gear",
		weight:     150,
		properties: []string{""},
		force: []Trait{
			{
				feature:    "bottom",
				id:         18,
				name:       "None",
				weight:     0,
				properties: []string{""},
			},
		},
	},
	{
		feature:    "top",
		id:         30,
		name:       "Farmer Apron",
		weight:     150,
		properties: []string{""},
		force: []Trait{
			{
				feature:    "bottom",
				id:         16,
				name:       "Farmer Apron",
				weight:     500,
				properties: []string{""},
			},
		},
	},
	{
		feature:    "top",
		id:         31,
		name:       "Thrifty Getup",
		weight:     150,
		properties: []string{""},
		force: []Trait{
			{
				feature:    "bottom",
				id:         17,
				name:       "Thrifty Getup",
				weight:     500,
				properties: []string{""},
			},
		},
	},
	{
		feature:    "top",
		id:         32,
		name:       "Red Chainmail",
		weight:     150,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         33,
		name:       "Blue Chainmail",
		weight:     150,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         34,
		name:       "Old School Uniform",
		weight:     150,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         35,
		name:       "Stars Hoodie",
		weight:     150,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         36,
		name:       "Kimono",
		weight:     150,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         37,
		name:       "Bowtie Jacket",
		weight:     150,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         38,
		name:       "Samurai Robe",
		weight:     150,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         39,
		name:       "Gold Chain",
		weight:     150,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         40,
		name:       "Starry Night Gown",
		weight:     75,
		properties: []string{""},
		force: []Trait{
			{
				feature:    "bottom",
				id:         18,
				name:       "None",
				weight:     0,
				properties: []string{""},
			},
		},
	},
	{
		feature:    "top",
		id:         41,
		name:       "Oversized Shirt",
		weight:     75,
		properties: []string{""},
		force: []Trait{
			{
				feature:    "bottom",
				id:         18,
				name:       "None",
				weight:     0,
				properties: []string{""},
			},
		},
	},
	{
		feature:    "top",
		id:         42,
		name:       "Skeleton Suit",
		weight:     75,
		properties: []string{""},
		force: []Trait{
			{
				feature:    "bottom",
				id:         18,
				name:       "None",
				weight:     0,
				properties: []string{""},
			},
		},
	},
	{
		feature:    "top",
		id:         43,
		name:       "Mummy Suit",
		weight:     75,
		properties: []string{""},
		force: []Trait{
			{
				feature:    "bottom",
				id:         18,
				name:       "None",
				weight:     0,
				properties: []string{""},
			},
		},
	},
	{
		feature:    "top",
		id:         44,
		name:       "Bathrobe",
		weight:     75,
		properties: []string{""},
		force: []Trait{
			{
				feature:    "bottom",
				id:         18,
				name:       "None",
				weight:     0,
				properties: []string{""},
			},
		},
	},
	{
		feature:    "top",
		id:         45,
		name:       "Traveler’s Armor",
		weight:     75,
		properties: []string{""},
		force: []Trait{
			{
				feature:    "bottom",
				id:         18,
				name:       "None",
				weight:     0,
				properties: []string{""},
			},
		},
	},
	{
		feature:    "top",
		id:         46,
		name:       "Bubbles",
		weight:     75,
		properties: []string{""},
		force: []Trait{
			{
				feature:    "bottom",
				id:         18,
				name:       "None",
				weight:     0,
				properties: []string{""},
			},
		},
	},
	{
		feature:    "top",
		id:         47,
		name:       "Future Suit",
		weight:     75,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         48,
		name:       "Maid Outfit",
		weight:     75,
		properties: []string{""},
		force: []Trait{
			{
				feature:    "bottom",
				id:         18,
				name:       "None",
				weight:     0,
				properties: []string{""},
			},
		},
	},
	{
		feature:    "top",
		id:         49,
		name:       "Bell Robe",
		weight:     75,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         50,
		name:       "Eagle Coat",
		weight:     75,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         51,
		name:       "Bomber Jacket",
		weight:     75,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         52,
		name:       "Gold Fancy",
		weight:     50,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         53,
		name:       "Edgelord Coat",
		weight:     50,
		properties: []string{""},
	},
	{
		feature:    "top",
		id:         54,
		name:       "Penguin Outfit",
		weight:     50,
		properties: []string{""},
		force: []Trait{
			{
				feature:    "bottom",
				id:         18,
				name:       "None",
				weight:     0,
				properties: []string{""},
			},
		},
	},
	{
		feature:    "top",
		id:         55,
		name:       "Admirals Coat",
		weight:     50,
		properties: []string{""},
	},
}

var fighterBodies = []Trait{
	{
		feature: "body",
		id:      0,
		name:    "Light",
		weight:  28,
	},
	{
		feature: "body",
		id:      1,
		name:    "Tan",
		weight:  28,
	},
	{
		feature: "body",
		id:      2,
		name:    "Clay",
		weight:  28,
	},
	{
		feature: "body",
		id:      3,
		name:    "Elf",
		weight:  7,
	},
	{
		feature: "body",
		id:      4,
		name:    "Ghostly",
		weight:  7,
	},
	{
		feature: "body",
		id:      5,
		name:    "Skeleton",
		weight:  1,
	},
}

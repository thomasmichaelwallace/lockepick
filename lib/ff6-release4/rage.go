package main

import "encoding/json"

type (
	Rage_type   uint8
	RageMastery [32]Rage_type
)

const (
	Guard Rage_type = iota
	ImperialSoldier
	Templar
	Ninja
	Samurai
	Borghese
	MagnaRoader_Purple
	Yojimbo
	Cloud
	Misty
	AlJabr
	Zaghrem
	Apocrypha
	DarkForce
	AngelWhisper
	Oversoul
	SkeletalHorror
	Commander
	Mu
	Wererat
	Mugbear
	Belmodar
	MuudSuud
	LeafBunny
	StrayCat
	SilverLobo
	Doberman
	Megalodoth
	Fidor
	Briareus
	Suriander
	Chimera
	Behemoth
	Fafnir
	LesserLopros
	FossilDragon
	HolyDragon
	FiendDragon
	Brachiosaur
	Tyrannosaur
	Darkwind
	Aepyornis
	Vulture
	Vasegiatta
	Zokka
	Trapper
	Hornet
	Nettlehopper
	DeltaBeetle
	KillerMantis
	Trillium
	Rafflesia
	Tumbleweed
	VampireThorn
	Cartagra
	Siegfried
	Nautiloid
	Exocite
	Anguiform
	LeapFrog
	Lizard
	LitworChicken
	Slagworm
	HellsRider
	Typhon
	OnionKnight
	MagitekArmor
	SkyArmor
	Satellite
	ArmoredWeapon
	Spritzer
	Flan
	Outcast
	Humpty
	Brainpan
	Cruller
	Cactuar_rage
	Bandit
	Harvester
	Bomb
	StillLife
	Lunatys
	VeilDancer
	HillGigas
	Tonberry
	MagicUrn
	Mover
	FigaroLizard
	Devoahan
	Aspiran
	Ghost_rage
	Crawler
	SandRay
	Alacran
	Actinian
	Sandhorse
	Darkside
	Malboro
	Urok
	Foper
	GuardLeader
	Corporal
	General
	Covert
	Kamui
	Warlock
	Cherry
	Joker
	IronFist
	Devil
	Provoker
	Cloudwraith
	Mahadeva
	VectorHound
	Peeper
	Stunner
	Sorath
	Destroyer
	Chippirabbit
	CoeurlCat
	Bloodfang
	HuntingHound
	Gorgias
	Don
	Murussu
	Wartpuck
	Gorgimera
	BehemothKing_Undead
	VectorLythos
	Wyvern
	ZombieDragon
	Dragon
	PrimevalDragon
	Weredragon
	Cirpius
	Sprinter
	Lenergia
	Marchosias
	Gloomwind
	Dropper
	RockWasp
	Grasswyrm
	Luridan
	Twinscythe
	Paraladia
	Exoray
	Crusher
	Ouroboros
	Acrophies
	Schmidt
	Devourer
	Cancer
	Gigantoad
	Basilisk
	MedusaChicken
	Landworm
	TestRider
	PlutoArmor
	OnionDasher
	HeavyArmor
	Chaser
	Gamma
	Poplium
	Intangir
	Misfit
	Creature
	Enuo
	Deepeye
	Unseelie
	NeckHunter
	Grenade
	AlluringRider
	Pandora
	BladeDancer
	Gigantos
	MagnaRoader_Red
	Lycaon
	Parasite
	LandRay
	Antares
	Anemone
	Moonform
	Specter
	GreatMalboro
	Bonnacon
	Oceanus
	LivingDead
	DeathWarden
	Face
	Outsider
	Coco
	Zeveak
	Nightwalker
	DemonKnight
	ImperialElite
	DesertHare
	Wizard
	DevilFist
	Illuyankas
	Sergeant
	Aspidochelon
	Knotty
	LunaWolf
	Belzecue
	Caladrius
	Tzakmaqiel
	Lukhavi
	Eukaryote
	LandGrillon
	Goetia
	GreaterMantis
	Bogy
	Purusa
	BlackDragon
	Adamankary
	Dante
	PlatinumDragon
	DuelArmor
	Psychos
	Mousse
	ShamblingCorpse
	Punisher_rage
	Balloon
	Gobbledygook
	GreatBehemoth
	Scorpion
	ChaosDragon
	Spitfire
	VectorChimera
	Lich
	Rukh
	MagnaRoader_Yellow
	Bug
	Seaflower
	Fortis
	Venobennu
	Galypdes
	Junk
	Mandrake
	Valeor
	Amduscias
	Necromancer
	GlasyaLabolas
	MagnaRoader_Brown
	WildRat
	GoldBear
	InnoSent
	Clymenus
	Garm
	Daedalus
	Baalzephon
	Ahriman
	DeathMachine
	MetalHitman
	Io
	Tonberries
)

var ragenames = map[Rage_type]string{
	Guard:               "Guard",
	ImperialSoldier:     "Imperial Soldier",
	Templar:             "Templar",
	Ninja:               "Ninja",
	Samurai:             "Samurai",
	Borghese:            "Borghese",
	MagnaRoader_Purple:  "Magna Roader (Purple)",
	Yojimbo:             "Yojimbo",
	Cloud:               "Cloud",
	Misty:               "Misty",
	AlJabr:              "Al Jabr",
	Zaghrem:             "Zaghrem",
	Apocrypha:           "Apocrypha",
	DarkForce:           "DarkForce",
	AngelWhisper:        "Angel Whisper",
	Oversoul:            "Oversoul",
	SkeletalHorror:      "Skeletal Horror",
	Commander:           "Commander",
	Mu:                  "Mu",
	Wererat:             "Wererat",
	Mugbear:             "Mugbear",
	Belmodar:            "Belmodar",
	MuudSuud:            "Muud Suud",
	LeafBunny:           "Leaf Bunny",
	StrayCat:            "Stray Cat",
	SilverLobo:          "Silver Lobo",
	Doberman:            "Doberman",
	Megalodoth:          "Megalodoth",
	Fidor:               "Fidor",
	Briareus:            "Briareus",
	Suriander:           "Suriander",
	Chimera:             "Chimera",
	Behemoth:            "Behemoth",
	Fafnir:              "Fafnir",
	LesserLopros:        "Lesser Lopros",
	FossilDragon:        "Fossil Dragon",
	HolyDragon:          "Holy Dragon",
	FiendDragon:         "Fiend Dragon",
	Brachiosaur:         "Brachiosaur",
	Tyrannosaur:         "Tyrannosaur",
	Darkwind:            "Darkwind",
	Aepyornis:           "Aepyornis",
	Vulture:             "Vulture",
	Vasegiatta:          "Vasegiatta",
	Zokka:               "Zokka",
	Trapper:             "Trapper",
	Hornet:              "Hornet",
	Nettlehopper:        "Nettlehopper",
	DeltaBeetle:         "Delta Beetle",
	KillerMantis:        "Killer Mantis",
	Trillium:            "Trillium",
	Rafflesia:           "Rafflesia",
	Tumbleweed:          "Tumbleweed",
	VampireThorn:        "Vampire Thorn",
	Cartagra:            "Cartagra",
	Siegfried:           "Siegfried",
	Nautiloid:           "Nautiloid",
	Exocite:             "Exocite",
	Anguiform:           "Anguiform",
	LeapFrog:            "Leap Frog",
	Lizard:              "Lizard",
	LitworChicken:       "Litwor Chicken",
	Slagworm:            "Slagworm",
	HellsRider:          "Hells Rider",
	Typhon:              "Typhon",
	OnionKnight:         "Onion Knight",
	MagitekArmor:        "Magitek Armor",
	SkyArmor:            "Sky Armor",
	Satellite:           "Satellite",
	ArmoredWeapon:       "Armored Weapon",
	Spritzer:            "Spritzer",
	Flan:                "Flan",
	Outcast:             "Outcast",
	Humpty:              "Humpty",
	Brainpan:            "Brainpan",
	Cruller:             "Cruller",
	Cactuar_rage:        "Cactuar",
	Bandit:              "Bandit",
	Harvester:           "Harvester",
	Bomb:                "Bomb",
	StillLife:           "Still Life",
	Lunatys:             "Lunatys",
	VeilDancer:          "Veil Dancer",
	HillGigas:           "Hill Gigas",
	Tonberry:            "Tonberry",
	MagicUrn:            "Magic Urn",
	Mover:               "Mover",
	FigaroLizard:        "Figaro Lizard",
	Devoahan:            "Devoahan",
	Aspiran:             "Aspiran",
	Ghost_rage:          "Ghost",
	Crawler:             "Crawler",
	SandRay:             "SandRay",
	Alacran:             "Alacran",
	Actinian:            "Actinian",
	Sandhorse:           "Sandhorse",
	Darkside:            "Darkside",
	Malboro:             "Malboro",
	Urok:                "Urok",
	Foper:               "Foper",
	GuardLeader:         "Guard Leader",
	Corporal:            "Corporal",
	General:             "General",
	Covert:              "Covert",
	Kamui:               "Kamui",
	Warlock:             "Warlock",
	Cherry:              "Cherry",
	Joker:               "Joker",
	IronFist:            "Iron Fist",
	Devil:               "Devil",
	Provoker:            "Provoker",
	Cloudwraith:         "Cloudwraith",
	Mahadeva:            "Mahadeva",
	VectorHound:         "Vector Hound",
	Peeper:              "Peeper",
	Stunner:             "Stunner",
	Sorath:              "Sorath",
	Destroyer:           "Destroyer",
	Chippirabbit:        "Chippirabbit",
	CoeurlCat:           "Coeurl Cat",
	Bloodfang:           "Bloodfang",
	HuntingHound:        "Hunting Hound",
	Gorgias:             "Gorgias",
	Don:                 "Don",
	Murussu:             "Murussu",
	Wartpuck:            "Wartpuck",
	Gorgimera:           "Gorgimera",
	BehemothKing_Undead: "Behemoth King (Undead)",
	VectorLythos:        "Vector Lythos",
	Wyvern:              "Wyvern",
	ZombieDragon:        "Zombie Dragon",
	Dragon:              "Dragon",
	PrimevalDragon:      "Primeval Dragon",
	Weredragon:          "Weredragon",
	Cirpius:             "Cirpius",
	Sprinter:            "Sprinter",
	Lenergia:            "Lenergia",
	Marchosias:          "Marchosias",
	Gloomwind:           "Gloomwind",
	Dropper:             "Dropper",
	RockWasp:            "Rock Wasp",
	Grasswyrm:           "Grasswyrm",
	Luridan:             "Luridan",
	Twinscythe:          "Twinscythe",
	Paraladia:           "Paraladia",
	Exoray:              "Exoray",
	Crusher:             "Crusher",
	Ouroboros:           "Ouroboros",
	Acrophies:           "Acrophies",
	Schmidt:             "Schmidt",
	Devourer:            "Devourer",
	Cancer:              "Cancer",
	Gigantoad:           "Gigantoad",
	Basilisk:            "Basilisk",
	MedusaChicken:       "MedusaChicken",
	Landworm:            "Landworm",
	TestRider:           "Test Rider",
	PlutoArmor:          "Pluto Armor",
	OnionDasher:         "Onion Dasher",
	HeavyArmor:          "Heavy Armor",
	Chaser:              "Chaser",
	Gamma:               "Gamma",
	Poplium:             "Poplium",
	Intangir:            "Intangir",
	Misfit:              "Misfit",
	Creature:            "Creature",
	Enuo:                "Enuo",
	Deepeye:             "Deepeye",
	Unseelie:            "Unseelie",
	NeckHunter:          "Neck Hunter",
	Grenade:             "Grenade",
	AlluringRider:       "Alluring Rider",
	Pandora:             "Pandora",
	BladeDancer:         "Blade Dancer",
	Gigantos:            "Gigantos",
	MagnaRoader_Red:     "Magna Roader (Red)",
	Lycaon:              "Lycaon",
	Parasite:            "Parasite",
	LandRay:             "Land Ray",
	Antares:             "Antares",
	Anemone:             "Anemone",
	Moonform:            "Moonform",
	Specter:             "Specter",
	GreatMalboro:        "Great Malboro",
	Bonnacon:            "Bonnacon",
	Oceanus:             "Oceanus",
	LivingDead:          "Living Dead",
	DeathWarden:         "Death Warden",
	Face:                "Face",
	Outsider:            "Outsider",
	Coco:                "Coco",
	Zeveak:              "Zeveak",
	Nightwalker:         "Nightwalker",
	DemonKnight:         "Demon Knight",
	ImperialElite:       "Imperial Elite",
	DesertHare:          "Desert Hare",
	Wizard:              "Wizard",
	DevilFist:           "Devil Fist",
	Illuyankas:          "Illuyankas",
	Sergeant:            "Sergeant",
	Aspidochelon:        "Aspidochelon",
	Knotty:              "Knotty",
	LunaWolf:            "Luna Wolf",
	Belzecue:            "Belzecue",
	Caladrius:           "Caladrius",
	Tzakmaqiel:          "Tzakmaqiel",
	Lukhavi:             "Lukhavi",
	Eukaryote:           "Eukaryote",
	LandGrillon:         "LandGrillon",
	Goetia:              "Goetia",
	GreaterMantis:       "Greater Mantis",
	Bogy:                "Bogy",
	Purusa:              "Purusa",
	BlackDragon:         "Black Dragon",
	Adamankary:          "Adamankary",
	Dante:               "Dante",
	PlatinumDragon:      "Platinum Dragon",
	DuelArmor:           "Duel Armor",
	Psychos:             "Psychos",
	Mousse:              "Mousse",
	ShamblingCorpse:     "Shambling Corpse",
	Punisher_rage:       "Punisher",
	Balloon:             "Balloon",
	Gobbledygook:        "Gobbledygook",
	GreatBehemoth:       "Great Behemoth",
	Scorpion:            "Scorpion",
	ChaosDragon:         "Chaos Dragon",
	Spitfire:            "Spitfire",
	VectorChimera:       "Vector Chimera",
	Lich:                "Lich",
	Rukh:                "Rukh",
	MagnaRoader_Yellow:  "Magna Roader (Yellow)",
	Bug:                 "Bug",
	Seaflower:           "Seaflower",
	Fortis:              "Fortis",
	Venobennu:           "Venobennu",
	Galypdes:            "Galypdes",
	Junk:                "Junk",
	Mandrake:            "Mandrake",
	Valeor:              "Valeor",
	Amduscias:           "Amduscias",
	Necromancer:         "Necromancer",
	GlasyaLabolas:       "Glasya Labolas",
	MagnaRoader_Brown:   "Magna Roader (Brown)",
	WildRat:             "Wild Rat",
	GoldBear:            "Gold Bear",
	InnoSent:            "Inno Sent",
	Clymenus:            "Clymenus",
	Garm:                "Garm",
	Daedalus:            "Daedalus",
	Baalzephon:          "Baalzephon",
	Ahriman:             "Ahriman",
	DeathMachine:        "Death Machine",
	MetalHitman:         "Metal Hitman",
	Io:                  "Io",
	Tonberries:          "Tonberries",
}

func (r RageMastery) MarshalJSON() ([]byte, error) {
	d := make(map[string]bool)
	for k, v := range ragenames {
		idx := k >> 3
		bit := (k & 7)

		d[v] = r[idx]&(1<<bit) != 0
	}
	return json.Marshal(d)
}

func (r *RageMastery) UnmarshalJSON(data []byte) error {
	d := make(map[string]bool)
	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}
	var rr RageMastery
	for k, v := range ragenames {
		idx := k >> 3
		bit := (k & 7)

		if d[v] {
			rr[idx] |= (1 << bit)
		}
	}
	*r = rr
	return nil
}
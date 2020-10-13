// Copyright 2014
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
)

type (
	Item      uint8
	Inventory struct {
		SortOrder [288]Item
		Count     [288]uint8
	}
	Equipable Item
)

const (

	// Dirk (0)
	Dagger Item = iota
	MythrilKnife
	MainGauche
	AirKnife
	ThiefsKnife
	AssassinsDagger
	ManEater
	Swordbreaker
	Gladius
	ValiantKnife
	// Sword (0xA)
	MythrilSword
	GreatSword
	RuneBlade
	Flametounge
	Icebrand
	ThunderBlade
	BastardSword
	Stoneblade
	BloodSword
	Enhancer
	CrystalSword
	Falchion
	SoulSabre
	Organyx
	Excalibur
	Zantetsuken
	Lightbringer
	Ragnarok
	UltimaWeapon
	// Lance (0x1d)
	MythrilSpear
	Trident
	HeavyLance
	Partisian
	HolyLance
	GoldenSpear
	RadiantLance
	Impartisian
	//  Dirk (0x25)
	Kunai
	Kodachi
	Sakura
	Sasuke
	Ichigeki
	Kagenui
	// Knife (0x2B)
	Ashura
	Kotetsu
	Kikuichimonji
	Kazekiri
	Murasame
	Masamune
	Murakuma
	Mutsunokami
	// Rod (0x33)
	HealingRod
	MythrilRod
	FlameRod
	IceRod
	ThunderRod
	PoisonRod
	HolyRod
	GravityRod
	Punisher
	MagusRod
	// Brush (0x3d)
	ChocoboBrush
	DaVinciBrush
	MagicalBrush
	RainbowBrush
	// Stars (0x41)
	Shuriken
	FumaShuriken
	Pinwheel
	// Special (0x44)
	ChainFlail
	MoonringBlade
	MorningStar
	Boomerang
	RisingSun
	Hawkeye
	BoneClub
	Sniper
	WingEdge
	// Gambler (0x4d)
	Cards
	Darts
	DeathTarot
	ViperDarts
	Dice
	FixedDice
	// Claw (0x53)
	MetalKnuckle
	MythrilClaw
	KaiserKnuckle
	VenomClaws
	BurningFist
	DragonClaws
	Tigerfang
	// Shields (0x5A)
	Buckler
	HeavyShield
	MythrilShield
	GoldenShield
	AegisShield
	DiamondShield
	FlameShield
	IceShield
	ThunderShield
	CrystalShield
	GenjiShield
	TortoiseShield
	CursedShield
	PaladinShield
	ForceShield
	// Helmet (0x69)
	LeatherHat
	Hairband
	PlumedHat
	Beret
	MagusHat
	Bandana
	IronHelmet
	HypnoCrown
	PriestsMiter
	GreenBeret
	TwistHeadband
	MythrilHelm
	Tiara
	GoldenHelmet
	TigerMask
	RedCap
	MysteryVeil
	Circlet
	RoyalCrown
	DiamondHelm
	BlackCowl
	CrystalHelm
	OathViel
	CatEarHood
	GenjiHelmet
	Thornlet
	Saucer
	// Armor (0x84)
	LeatherArmor
	CottonRobe
	KenpoGi
	IronArmor
	SilkRobe
	MythrilVest
	NinjaGear
	WhiteDress
	MythrilMail
	GaiaGear
	MirageDress
	GoldenArmor
	PowerSash
	LuminousRobe
	DiamondVest
	RedJacket
	ForceArmor
	DiamondArmor
	BlackGarb
	MagusRobe
	CrystalMail
	RegalGown
	GenjiArmor
	ReedCloak
	MinervaBustier
	TabbySuit
	ChocoboSuit
	MoogleSuit
	NutkinSuit
	BehemothSuit
	SnowScarf
	// Tool (0xa3)
	Noiseblaster
	Bioblaster
	Flash
	Chainsaw
	Debilitator
	Drill
	AirAnchor
	AutoCrossbow
	// Skean (0xAB)
	FlameScroll
	WaterScroll
	LightningScroll
	InvisibilityScroll
	ShadowScroll
	// Relics (0xB0)
	SilverSpectacles
	StarPendant
	PeaceRing
	Amulet
	WhiteCape
	JeweledRing
	FairyRing
	BarrierRing
	MythrilGlove
	ProtectRing
	HermesSandals
	ReflectRing
	AngelWings
	AngelRing
	KnightsCode
	DragoonBoots
	ZephyrCloak
	PrincessRing
	CursedRing
	Earrings
	GigasGlove
	BlizzardRing
	BerserkerRing
	ThiefsBracer
	GuardBracelet
	HerosRing
	Ribbon
	MuscleBelt
	CrystalOrb
	GoldHairpin
	Celestriad
	BrigandsGlove
	Gauntlet
	GenjiGlove
	HyperWrist
	MastersScroll
	PrayerBeads
	BlackBelt
	HejisJitte
	FakeMoustache
	SoulOfTamasa
	DragonHorn
	MeritAward
	MementoRing
	SafetyBit
	LichRing
	MolulusCharm
	WardBangle
	MiracleShoes
	AlarmEarring
	GaleHairpin
	SniperEye
	GrowthEgg
	Tintinabulum
	SprintShoes
	// Misc (0xE7)
	RenameCard
	Potion
	HiPotion
	XPotion
	Ether
	HiEther
	XEther
	Elixir
	Megalixir
	PhoenixDown
	HolyWater
	Antidote
	EyeDrops
	GoldNeedle
	Remedy
	SleepingBag
	Tent
	GreenCherry
	MagiciteShard
	SuperBall
	EchoScreen
	SmokeBomb
	TeleportStone
	DriedMeat
	Nothing
)
const (
	// Odd huh? :)
	Apocalypse = Equipable(RenameCard) + iota
	ZwillCrossblade
	Zanmato
	Oborozuki
	Longinus
	Godhand
	SavetheQueen
	StardustRod
	AngelBrush
	FinalTrump
	Gungnir
	DuelingMask
	ScorpionTail
	BoneWrist
	Excalipoor
)

var (
	iname = map[Item]string{
		Dagger:             "Dagger",
		MythrilKnife:       "Mythril Knife",
		MainGauche:         "Main Gauche",
		AirKnife:           "Air Knife",
		ThiefsKnife:        "Thief's Knife",
		AssassinsDagger:    "Assassin's Dagger",
		ManEater:           "Man-Eater",
		Swordbreaker:       "Swordbreaker",
		Gladius:            "Gladius",
		ValiantKnife:       "Valiant Knife",
		MythrilSword:       "Mythril Sword",
		GreatSword:         "Great Sword",
		RuneBlade:          "Rune Blade",
		Flametounge:        "Flametounge",
		Icebrand:           "Icebrand",
		ThunderBlade:       "Thunder Blade",
		BastardSword:       "Bastard Sword",
		Stoneblade:         "Stoneblade",
		BloodSword:         "Blood Sword",
		Enhancer:           "Enhancer",
		CrystalSword:       "Crystal Sword",
		Falchion:           "Falchion",
		SoulSabre:          "Soul Sabre",
		Organyx:            "Organyx",
		Excalibur:          "Excalibur",
		Zantetsuken:        "Zantetsuken",
		Lightbringer:       "Lightbringer",
		Ragnarok:           "Ragnarok",
		UltimaWeapon:       "Ultima Weapon",
		MythrilSpear:       "Mythril Spear",
		Trident:            "Trident",
		HeavyLance:         "Heavy Lance",
		Partisian:          "Partisian",
		HolyLance:          "Holy Lance",
		GoldenSpear:        "Golden Spear",
		RadiantLance:       "Radiant Lance",
		Impartisian:        "Impartisian",
		Kunai:              "Kunai",
		Kodachi:            "Kodachi",
		Sakura:             "Sakura",
		Sasuke:             "Sasuke",
		Ichigeki:           "Ichigeki",
		Kagenui:            "Kagenui",
		Ashura:             "Ashura",
		Kotetsu:            "Kotetsu",
		Kikuichimonji:      "Kiku-ichimonji",
		Kazekiri:           "Kazekiri",
		Murasame:           "Murasame",
		Masamune:           "Masamune",
		Murakuma:           "Murakuma",
		Mutsunokami:        "Mutsunokami",
		HealingRod:         "Healing Rod",
		MythrilRod:         "Mythril Rod",
		FlameRod:           "Flame Rod",
		IceRod:             "Ice Rod",
		ThunderRod:         "Thunder Rod",
		PoisonRod:          "Poison Rod",
		HolyRod:            "Holy Rod",
		GravityRod:         "Gravity Rod",
		Punisher:           "Punisher",
		MagusRod:           "Magus Rod",
		ChocoboBrush:       "Chocobo Brush",
		DaVinciBrush:       "DaVinci Brush",
		MagicalBrush:       "Magical Brush",
		RainbowBrush:       "Rainbow Brush",
		Shuriken:           "Shuriken",
		FumaShuriken:       "Fuma Shuriken",
		Pinwheel:           "Pinwheel",
		ChainFlail:         "Chain Flail",
		MoonringBlade:      "Moonring Blade",
		MorningStar:        "Morning Star",
		Boomerang:          "Boomerang",
		RisingSun:          "Rising Sun",
		Hawkeye:            "Hawkeye",
		BoneClub:           "Bone Club",
		Sniper:             "Sniper",
		WingEdge:           "Wing Edge",
		Cards:              "Cards",
		Darts:              "Darts",
		DeathTarot:         "Death Tarot",
		ViperDarts:         "Viper Darts",
		Dice:               "Dice",
		FixedDice:          "Fixed Dice",
		MetalKnuckle:       "Metal Knuckle",
		MythrilClaw:        "Mythril Claw",
		KaiserKnuckle:      "Kaiser Knuckle",
		VenomClaws:         "Venom Claws",
		BurningFist:        "Burning Fist",
		DragonClaws:        "Dragon Claws",
		Tigerfang:          "Tigerfang",
		Buckler:            "Buckler",
		HeavyShield:        "Heavy Shield",
		MythrilShield:      "Mythril Shield",
		GoldenShield:       "Golden Shield",
		AegisShield:        "Aegis Shield",
		DiamondShield:      "Diamond Shield",
		FlameShield:        "Flame Shield",
		IceShield:          "Ice Shield",
		ThunderShield:      "Thunder Shield",
		CrystalShield:      "Crystal Shield",
		GenjiShield:        "Genji Shield",
		TortoiseShield:     "Tortoise Shield",
		CursedShield:       "Cursed Shield",
		PaladinShield:      "Paladin's Shield",
		ForceShield:        "Force Shield",
		LeatherHat:         "Leather Hat",
		Hairband:           "Hairband",
		PlumedHat:          "Plumed Hat",
		Beret:              "Beret",
		MagusHat:           "Magus Hat",
		Bandana:            "Bandana",
		IronHelmet:         "Iron Helmet",
		HypnoCrown:         "Hypno Crown",
		PriestsMiter:       "Priest's Miter",
		GreenBeret:         "Green Beret",
		TwistHeadband:      "Twist Headband",
		MythrilHelm:        "Mythril Helm",
		Tiara:              "Tiara",
		GoldenHelmet:       "Golden Helmet",
		TigerMask:          "Tiger Mask",
		RedCap:             "Red Cap",
		MysteryVeil:        "Mystery Veil",
		Circlet:            "Circlet",
		RoyalCrown:         "Royal Crown",
		DiamondHelm:        "Diamond Helm",
		BlackCowl:          "Black Cowl",
		CrystalHelm:        "Crystal Helm",
		OathViel:           "Oath Viel",
		CatEarHood:         "Cat-Ear Hood",
		GenjiHelmet:        "Genji Helmet",
		Thornlet:           "Thornlet",
		Saucer:             "Saucer",
		LeatherArmor:       "Leather Armor",
		CottonRobe:         "Cotton Robe",
		KenpoGi:            "Kenpo Gi",
		IronArmor:          "Iron Armor",
		SilkRobe:           "Silk Robe",
		MythrilVest:        "Mythril Vest",
		NinjaGear:          "Ninja Gear",
		WhiteDress:         "White Dress",
		MythrilMail:        "Mythril Mail",
		GaiaGear:           "Gaia Gear",
		MirageDress:        "Mirage Dress",
		GoldenArmor:        "Golden Armor",
		PowerSash:          "Power Sash",
		LuminousRobe:       "Luminous Robe",
		DiamondVest:        "Diamond Vest",
		RedJacket:          "Red Jacket",
		ForceArmor:         "Force Armor",
		DiamondArmor:       "Diamond Armor",
		BlackGarb:          "Black Garb",
		MagusRobe:          "Magus Robe",
		CrystalMail:        "Crystal Mail",
		RegalGown:          "Regal Gown",
		GenjiArmor:         "Genji Armor",
		ReedCloak:          "Reed Cloak",
		MinervaBustier:     "Minerva Bustier",
		TabbySuit:          "Tabby Suit",
		ChocoboSuit:        "Chocobo Suit",
		MoogleSuit:         "Moogle Suit",
		NutkinSuit:         "Nutkin Suit",
		BehemothSuit:       "Behemoth Suit",
		SnowScarf:          "Snow Scarf",
		Noiseblaster:       "Noiseblaster",
		Bioblaster:         "Bioblaster",
		Flash:              "Flash",
		Chainsaw:           "Chainsaw",
		Debilitator:        "Debilitator",
		Drill:              "Drill",
		AirAnchor:          "Air Anchor",
		AutoCrossbow:       "Auto Crossbow",
		FlameScroll:        "Flame Scroll",
		WaterScroll:        "Water Scroll",
		LightningScroll:    "Lightning Scroll",
		InvisibilityScroll: "Invisibility Scroll",
		ShadowScroll:       "Shadow Scroll",
		SilverSpectacles:   "Silver Spectacles",
		StarPendant:        "Star Pendant",
		PeaceRing:          "Peace Ring",
		Amulet:             "Amulet",
		WhiteCape:          "White Cape",
		JeweledRing:        "Jeweled Ring",
		FairyRing:          "Fairy Ring",
		BarrierRing:        "Barrier Ring",
		MythrilGlove:       "Mythril Glove",
		ProtectRing:        "Protect Ring",
		HermesSandals:      "Hermes Sandals",
		ReflectRing:        "Reflect Ring",
		AngelWings:         "Angel Wings",
		AngelRing:          "Angel Ring",
		KnightsCode:        "Knights Code",
		DragoonBoots:       "Dragon Boots",
		ZephyrCloak:        "Zephyr Cloak",
		PrincessRing:       "Princess Ring",
		CursedRing:         "Cursed Ring",
		Earrings:           "Earrings",
		GigasGlove:         "Gigas Glove",
		BlizzardRing:       "Blizzard Ring",
		BerserkerRing:      "Berserker Ring",
		ThiefsBracer:       "Thief's Bracer",
		GuardBracelet:      "Guard Bracelet",
		HerosRing:          "Hero's Ring",
		Ribbon:             "Ribbon",
		MuscleBelt:         "Muscle Belt",
		CrystalOrb:         "Crystal Orb",
		GoldHairpin:        "Gold Hairpin",
		Celestriad:         "Celestriad",
		BrigandsGlove:      "Brigand's Glove",
		Gauntlet:           "Gauntlet",
		GenjiGlove:         "Genji Glove",
		HyperWrist:         "Hyper Wrist",
		MastersScroll:      "Master's Scroll",
		PrayerBeads:        "Prayer Beads",
		BlackBelt:          "Black Belt",
		HejisJitte:         "Heji's Jitte",
		FakeMoustache:      "Fake Moustache",
		SoulOfTamasa:       "Soul Of Tamasa",
		DragonHorn:         "Dragon Horn",
		MeritAward:         "Merit Award",
		MementoRing:        "Memento Ring",
		SafetyBit:          "Safety Bit",
		LichRing:           "Lich Ring",
		MolulusCharm:       "Molulu's Charm",
		WardBangle:         "Ward Bangle",
		MiracleShoes:       "Miracle Shoes",
		AlarmEarring:       "Alarm Earring",
		GaleHairpin:        "Gale Hairpin",
		SniperEye:          "Sniper Eye",
		GrowthEgg:          "Growth Egg",
		Tintinabulum:       "Tintinabulum",
		SprintShoes:        "Sprint Shoes",
		RenameCard:         "Rename Card",
		Potion:             "Potion",
		HiPotion:           "Hi-Potion",
		XPotion:            "X-Potion",
		Ether:              "Ether",
		HiEther:            "Hi-Ether",
		XEther:             "X-Ether",
		Elixir:             "Elixir",
		Megalixir:          "Megalixir",
		PhoenixDown:        "Phoenix Down",
		HolyWater:          "Holy Water",
		Antidote:           "Antidote",
		EyeDrops:           "Eye Dropss",
		GoldNeedle:         "Gold Needle",
		Remedy:             "Remedy",
		SleepingBag:        "Sleeping Bag",
		Tent:               "Tent",
		GreenCherry:        "Green Cherry",
		MagiciteShard:      "Magicite Shard",
		SuperBall:          "Super Ball",
		EchoScreen:         "Echo Screen",
		SmokeBomb:          "Smoke Bomb",
		TeleportStone:      "Teleport Stone",
		DriedMeat:          "Dried Meat",
		Nothing:            "Nothing",
	}
	eqname = map[Equipable]string{
		Apocalypse:      "Apocalypse",
		ZwillCrossblade: "Zwill Crossblade",
		Zanmato:         "Zanmato",
		Oborozuki:       "Oborozuki",
		Longinus:        "Longinus",
		Godhand:         "Godhand",
		SavetheQueen:    "Save the Queen",
		StardustRod:     "Stardust Rod",
		AngelBrush:      "Angel Brush",
		FinalTrump:      "Final Trump",
		Gungnir:         "Gungnir",
		DuelingMask:     "Dueling Mask",
		ScorpionTail:    "Scorpion Tail",
		BoneWrist:       "Bone Wrist",
		Excalipoor:      "Excalipoor",
	}
)

func (e Equipable) MarshalJSON() ([]byte, error) {
	v, ok := eqname[e]
	if !ok {
		v = iname[Item(e)]
	}
	return json.Marshal(v)
}

func (e *Equipable) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	for k, v := range eqname {
		if v == s {
			*e = k
			return nil
		}
	}
	return (*Item)(e).UnmarshalJSON(data)
}

func (i Item) MarshalJSON() ([]byte, error) {
	return json.Marshal(iname[i])
}

func (i *Item) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	for k, v := range iname {
		if v == s {
			*i = k
			return nil
		}
	}
	return fmt.Errorf("\"%s\" isn't a known Item", s)
}

func (inv *Inventory) UnmarshalJSON(data []byte) error {
	d := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}

	if err := json.Unmarshal(d["SortingOrder"], &inv.SortOrder); err != nil {
		return err
	}
	for i, k := range inv.SortOrder {
		if err := json.Unmarshal(d[iname[k]], &inv.Count[i]); err != nil {
			return err
		}
	}
	for i := 256; i < len(inv.Count); i++ {
		inv.Count[i] = 0
		inv.SortOrder[i] = 0xff
	}
	return nil
}

func (inv Inventory) MarshalJSON() ([]byte, error) {
	d := make(map[string]interface{})
	d["SortingOrder"] = &inv.SortOrder
	for i, k := range inv.SortOrder {
		d[iname[k]] = inv.Count[i]
	}
	return json.Marshal(d)
}

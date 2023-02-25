import{A as t,cm as n,cn as e,c1 as a,Q as i,n as s,o as l,co as o,cp as r,cq as d,cr as m,cs as h,ct as c,cu as u,L as p,bL as g,bG as f,bJ as S,bt as b,T as y,d as v,a as T,ao as w,a$ as M,al as C,K as O,ap as W}from"./detailed_results-44ea1c48.chunk.js";import{m as I,u as L,a as N,b as P,c as F,B as R,I as k,v as A,w as B,T as D}from"./individual_sim_ui-563b8f34.chunk.js";import{T as E}from"./totem_inputs-84ddffd1.chunk.js";const G=I({fieldName:"bloodlust",id:t.fromSpellId(2825)}),j=L({fieldName:"shield",values:[{value:n.NoShield,tooltip:"No Shield"},{actionId:t.fromSpellId(57960),value:n.WaterShield},{actionId:t.fromSpellId(49281),value:n.LightningShield}]}),x={inputs:[N({fieldName:"type",label:"Type",values:[{name:"Adaptive",value:e.Adaptive,tooltip:"Dynamically adapts based on available mana to maximize CL casts without going OOM."},{name:"Manual",value:e.Manual,tooltip:"Allows custom selection of which spells to use and to modify cast conditions."}]}),P({fieldName:"inThunderstormRange",label:"In Thunderstorm Range",labelTooltip:"Thunderstorm will hit all targets when cast. Ignores knockback.",showWhen:t=>t.getTalents().thunderstorm}),F({fieldName:"lvbFsWaitMs",label:"Max wait for LvB/FS (ms)",labelTooltip:"Amount of time the sim will wait if FS is about to fall off or LvB CD is about to come up. Setting to 0 will default to 175ms"}),P({fieldName:"useChainLightning",label:"Use Chain Lightning in Rotation",labelTooltip:"Use Chain Lightning in rotation",enableWhen:t=>t.getRotation().type==e.Manual}),P({fieldName:"useClOnlyGap",label:"Use CL only as gap filler",labelTooltip:"Use CL to fill short gaps in LvB CD instead of on CD.",enableWhen:t=>t.getRotation().type==e.Manual&&t.getRotation().useChainLightning}),F({fieldName:"clMinManaPer",label:"Min mana percent to use Chain Lightning",labelTooltip:"Customize minimum mana level to cast Chain Lightning. 0 will spam until OOM.",enableWhen:t=>t.getRotation().type==e.Manual&&t.getRotation().useChainLightning}),P({fieldName:"useFireNova",label:"Use Fire Nova in Rotation",labelTooltip:"Fire Nova will hit all targets when cast.",enableWhen:t=>t.getRotation().type==e.Manual}),F({fieldName:"fnMinManaPer",label:"Min mana percent to use FireNova",labelTooltip:"Customize minimum mana level to cast Fire Nova. 0 will spam until OOM.",enableWhen:t=>t.getRotation().type==e.Manual&&t.getRotation().useFireNova}),P({fieldName:"overwriteFlameshock",label:"Allow Flameshock to be overwritten",labelTooltip:"Will use flameshock at the end of the duration even if its still ticking if there isn't enough time to cast lavaburst before expiring.",enableWhen:t=>t.getRotation().type==e.Manual}),P({fieldName:"alwaysCritLvb",label:"Only cast Lavaburst with FS",labelTooltip:"Will only cast Lavaburst if Flameshock will be active when the cast finishes.",enableWhen:t=>t.getRotation().type==e.Manual}),P({fieldName:"useThunderstorm",label:"Allow Thunderstorm to be cast.",labelTooltip:"Disabling this will stop thunderstorm from being cast entirely.",enableWhen:t=>t.getRotation().type==e.Manual,showWhen:t=>t.getTalents().thunderstorm})]},H={name:"Standard",data:a.create({talentsString:"0532001523212351322301351-005052031",glyphs:i.create({major1:s.GlyphOfLava,major2:s.GlyphOfTotemOfWrath,major3:s.GlyphOfLightningBolt,minor1:l.GlyphOfThunderstorm,minor2:l.GlyphOfWaterShield,minor3:l.GlyphOfGhostWolf})})},U=o.create({totems:r.create({earth:d.StrengthOfEarthTotem,air:m.WrathOfAirTotem,fire:h.TotemOfWrath,water:c.ManaSpringTotem}),type:e.Adaptive,fnMinManaPer:66,clMinManaPer:33,useChainLightning:!1,useFireNova:!1,useThunderstorm:!0}),J=u.create({shield:n.WaterShield,bloodlust:!0}),_=p.create({defaultPotion:g.RunicManaInjector,flask:f.FlaskOfTheFrostWyrm,food:S.FoodFishFeast}),z={name:"Pre-raid Preset",tooltip:R,gear:b.fromJsonString('{"items": [\n\t\t{\n\t\t\t"id": 37180,\n\t\t\t"enchant": 3820,\n\t\t\t"gems": [\n\t\t\t\t41285,\n\t\t\t\t42144\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 37595\n\t\t},\n\t\t{\n\t\t\t"id": 37673,\n\t\t\t"enchant": 3810,\n\t\t\t"gems": [\n\t\t\t\t42144\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 41610,\n\t\t\t"enchant": 3722\n\t\t},\n\t\t{\n\t\t\t"id": 39592,\n\t\t\t"enchant": 3832,\n\t\t\t"gems": [\n\t\t\t\t42144,\n\t\t\t\t40025\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 37788,\n\t\t\t"enchant": 2332,\n\t\t\t"gems": [\n\t\t\t\t0\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 39593,\n\t\t\t"enchant": 3246,\n\t\t\t"gems": [\n\t\t\t\t40051,\n\t\t\t\t0\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 40696,\n\t\t\t"gems": [\n\t\t\t\t40049,\n\t\t\t\t39998\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 37791,\n\t\t\t"enchant": 3719\n\t\t},\n\t\t{\n\t\t\t"id": 44202,\n\t\t\t"enchant": 3826,\n\t\t\t"gems": [\n\t\t\t\t39998\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 43253,\n\t\t\t"gems": [\n\t\t\t\t40027\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 37694\n\t\t},\n\t\t{\n\t\t\t"id": 40682\n\t\t},\n\t\t{\n\t\t\t"id": 37873\n\t\t},\n\t\t{\n\t\t\t"id": 41384,\n\t\t\t"enchant": 3834\n\t\t},\n\t\t{\n\t\t\t"id": 40698\n\t\t},\n\t\t{\n\t\t\t"id": 40708\n\t\t}\n  ]}')},q={name:"P1 Preset",tooltip:R,gear:b.fromJsonString('{"items": [\n\t\t{\n\t\t\t"id": 40516,\n\t\t\t"enchant": 3820,\n\t\t\t"gems": [\n\t\t\t\t41285,\n\t\t\t\t40027\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 44661,\n\t\t\t"gems": [\n\t\t\t\t39998\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 40286,\n\t\t\t"enchant": 3810\n\t\t},\n\t\t{\n\t\t\t"id": 44005,\n\t\t\t"enchant": 3722,\n\t\t\t"gems": [\n\t\t\t\t40027\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 40514,\n\t\t\t"enchant": 3832,\n\t\t\t"gems": [\n\t\t\t\t42144,\n\t\t\t\t42144\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 40324,\n\t\t\t"enchant": 2332,\n\t\t\t"gems": [\n\t\t\t\t42144,\n\t\t\t\t0\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 40302,\n\t\t\t"enchant": 3246,\n\t\t\t"gems": [\n\t\t\t\t0\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 40301,\n\t\t\t"gems": [\n\t\t\t\t40014\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 40560,\n\t\t\t"enchant": 3721\n\t\t},\n\t\t{\n\t\t\t"id": 40519,\n\t\t\t"enchant": 3826\n\t\t},\n\t\t{\n\t\t\t"id": 37694\n\t\t},\n\t\t{\n\t\t\t"id": 40399\n\t\t},\n\t\t{\n\t\t\t"id": 40432\n\t\t},\n\t\t{\n\t\t\t"id": 40255\n\t\t},\n\t\t{\n\t\t\t"id": 40395,\n\t\t\t"enchant": 3834\n\t\t},\n\t\t{\n\t\t\t"id": 40401,\n\t\t\t"enchant": 1128\n\t\t},\n\t\t{\n\t\t\t"id": 40267\n\t\t}\n  ]}')},K={name:"P2 Preset",tooltip:R,gear:b.fromJsonString('{"items": [\n        {\n          "id": 46209,\n          "enchant": 3820,\n          "gems": [\n            41285,\n            40048\n          ]\n        },\n        {\n          "id": 45933,\n          "gems": [\n            39998\n          ]\n        },\n        {\n          "id": 46211,\n          "enchant": 3810,\n          "gems": [\n            39998\n          ]\n        },\n        {\n          "id": 45242,\n          "enchant": 3722,\n          "gems": [\n            39998\n          ]\n        },\n        {\n          "id": 46206,\n          "enchant": 3832,\n          "gems": [\n            39998,\n            39998\n          ]\n        },\n        {\n          "id": 45460,\n          "enchant": 2332,\n          "gems": [\n            39998,\n            0\n          ]\n        },\n        {\n          "id": 45665,\n          "enchant": 3604,\n          "gems": [\n            39998,\n            39998,\n            0\n          ]\n        },\n        {\n          "id": 45616,\n          "enchant": 3599,\n          "gems": [\n            39998,\n            39998,\n            39998\n          ]\n        },\n        {\n          "id": 46210,\n          "enchant": 3721,\n          "gems": [\n            39998,\n            40027\n          ]\n        },\n        {\n          "id": 45537,\n          "enchant": 3606,\n          "gems": [\n            39998,\n            40027\n          ]\n        },\n        {\n          "id": 46046,\n          "gems": [\n            39998\n          ]\n        },\n        {\n          "id": 45495,\n          "gems": [\n            39998\n          ]\n        },\n        {\n          "id": 45518\n        },\n        {\n          "id": 40255\n        },\n        {\n          "id": 45612,\n          "enchant": 3834,\n          "gems": [\n            39998\n          ]\n        },\n        {\n          "id": 45470,\n          "enchant": 1128,\n          "gems": [\n            39998\n          ]\n        },\n        {\n          "id": 40267\n        }\n      ]}')};class Q extends k{constructor(t,n){super(t,n,{cssClass:"elemental-shaman-sim-ui",cssScheme:"shaman",knownIssues:[],warnings:[t=>({updateOn:y.onAny([t.player.rotationChangeEmitter,t.player.currentStatsEmitter]),getContent:()=>{const n=t.player.getCurrentStats().sets.includes("Skyshatter Regalia (2pc)"),e=t.player.getRotation().totems,a=e&&e.earth&&e.air&&e.fire&&e.water;return n&&!a?"T6 2pc bonus is equipped, but inactive because not all 4 totem types are being used.":""}})],epStats:[v.StatIntellect,v.StatSpellPower,v.StatSpellHit,v.StatSpellCrit,v.StatSpellHaste,v.StatMP5],epReferenceStat:v.StatSpellPower,displayStats:[v.StatHealth,v.StatMana,v.StatStamina,v.StatIntellect,v.StatSpellPower,v.StatSpellHit,v.StatSpellCrit,v.StatSpellHaste,v.StatMP5],modifyDisplayStats:t=>{let n=new T;return n=n.addStat(v.StatSpellHit,t.getTalents().elementalPrecision*A),n=n.addStat(v.StatSpellCrit,1*t.getTalents().tidalMastery*B),{talents:n}},defaults:{gear:q.gear,epWeights:T.fromMap({[v.StatIntellect]:.22,[v.StatSpellPower]:1,[v.StatSpellCrit]:.67,[v.StatSpellHaste]:1.29,[v.StatMP5]:.08}),consumes:_,rotation:U,talents:H.data,specOptions:J,raidBuffs:w.create({arcaneBrilliance:!0,divineSpirit:!0,giftOfTheWild:M.TristateEffectImproved,moonkinAura:M.TristateEffectImproved,sanctifiedRetribution:!0}),partyBuffs:C.create({}),individualBuffs:O.create({blessingOfKings:!0,blessingOfWisdom:2,vampiricTouch:!0}),debuffs:W.create({faerieFire:M.TristateEffectImproved,judgementOfWisdom:!0,misery:!0,curseOfElements:!0,shadowMastery:!0})},playerIconInputs:[j,G],rotationInputs:x,includeBuffDebuffInputs:[],excludeBuffDebuffInputs:[],otherInputs:{inputs:[D]},customSections:[E],encounterPicker:{showExecuteProportion:!1},presets:{talents:[H],gear:[z,q,K]}})}}export{U as D,Q as E,q as P,H as S,J as a,_ as b,K as c};
//# sourceMappingURL=sim-9bf2c13d.chunk.js.map

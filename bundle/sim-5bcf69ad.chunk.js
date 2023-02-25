import{A as t,aq as n,ar as e,c4 as a,c5 as s,c6 as i,c1 as o,Q as l,D as r,e as d,c7 as m,c8 as c,L as p,bL as h,bG as u,bJ as f,ao as g,a$ as S,K as b,al as y,ap as T,bt as w,d as v,a as I}from"./detailed_results-44ea1c48.chunk.js";import{m as M,a as O,b as P,c as W,B,I as C,M as R,d as x,A as N,e as U,f as F,g as A,h as D,T as E,D as H}from"./individual_sim_ui-563b8f34.chunk.js";const k=M({fieldName:"innervateTarget",id:t.fromSpellId(29166),extraCssClasses:["within-raid-sim-hide"],getValue:t=>t.getSpecOptions().innervateTarget?.targetIndex!=n,setValue:(t,a,s)=>{const i=a.getSpecOptions();i.innervateTarget=e.create({targetIndex:s?0:n}),a.setSpecOptions(t,i)}}),j={inputs:[O({fieldName:"type",label:"Type",labelTooltip:"If set to 'Adaptive', will dynamically adjust rotation.",values:[{name:"Adaptive",value:a.Adaptive,tooltip:"Dynamically adapts based on available mana to maximize CL casts without going OOM."},{name:"Manual",value:a.Manual,tooltip:"Allows custom selection of which spells to use, dot managment and CD usage."}]}),P({fieldName:"useSmartCooldowns",label:"Smart Cooldowns usage",labelTooltip:"The rotation will use cooldowns during eclipses, avoiding Haste CDs in solar and Crit CDs in lunar",showWhen:t=>t.getRotation().type==a.Manual}),O({fieldName:"mfUsage",label:"Moonfire Usage",labelTooltip:"Defines how Moonfire will be used in the rotation.",values:[{name:"Unused",value:s.NoMf},{name:"Before lunar",value:s.BeforeLunar},{name:"Maximize",value:s.MaximizeMf}],showWhen:t=>t.getRotation().type==a.Manual}),O({fieldName:"isUsage",label:"Insect Swarm Usage",labelTooltip:"Defines how Insect Swarm will be used in the rotation.",values:[{name:"Unused",value:i.NoIs},{name:"Before solar",value:i.BeforeSolar},{name:"Maximize",value:i.MaximizeIs}],showWhen:t=>t.getRotation().type==a.Manual}),P({fieldName:"useStarfire",label:"Use Starfire",labelTooltip:"Should the rotation use Starfire.",showWhen:t=>t.getRotation().type==a.Manual}),P({fieldName:"useWrath",label:"Use Wrath",labelTooltip:"Should the rotation use Wrath.",showWhen:t=>t.getRotation().type==a.Manual}),P({fieldName:"useTyphoon",label:"Use Typhoon",labelTooltip:"Should the rotation use Typhoon.",showWhen:t=>t.getRotation().type==a.Manual}),P({fieldName:"useHurricane",label:"Use Hurricane",labelTooltip:"Should the rotation use Hurricane.",showWhen:t=>t.getRotation().type==a.Manual}),P({fieldName:"useBattleRes",label:"Use Battle Res",labelTooltip:"Cast Battle Res on an ally sometime during the encounter.",showWhen:t=>t.getRotation().type==a.Manual}),W({fieldName:"playerLatency",label:"Player latency",labelTooltip:"Time before the player reacts to an eclipse proc, in milliseconds.",showWhen:t=>t.getRotation().type==a.Manual})]},z={name:"Phase 1",data:o.create({talentsString:"5032003115331303213305311231--205003012",glyphs:l.create({major1:r.GlyphOfFocus,major2:r.GlyphOfInsectSwarm,major3:r.GlyphOfStarfall,minor1:d.GlyphOfTyphoon,minor2:d.GlyphOfUnburdenedRebirth,minor3:d.GlyphOfTheWild})})},G=m.create({type:a.Adaptive,maintainFaerieFire:!0,useSmartCooldowns:!0,mfUsage:s.NoMf,isUsage:i.MaximizeIs,useStarfire:!0,useWrath:!0,useBattleRes:!1,playerLatency:200}),L=c.create({innervateTarget:e.create({targetIndex:n})}),J=p.create({defaultPotion:h.PotionOfSpeed,flask:u.FlaskOfTheFrostWyrm,food:f.FoodFishFeast,prepopPotion:h.PotionOfWildMagic,thermalSapper:!0}),_=g.create({arcaneBrilliance:!0,bloodlust:!0,divineSpirit:!0,giftOfTheWild:S.TristateEffectImproved,icyTalons:!0,moonkinAura:S.TristateEffectImproved,leaderOfThePack:S.TristateEffectImproved,powerWordFortitude:S.TristateEffectImproved,sanctifiedRetribution:!0,strengthOfEarthTotem:S.TristateEffectImproved,trueshotAura:!0,wrathOfAirTotem:!0}),K=b.create({blessingOfKings:!0,blessingOfMight:S.TristateEffectImproved,blessingOfWisdom:S.TristateEffectImproved,vampiricTouch:!0}),V=y.create({heroicPresence:!1}),q=T.create({bloodFrenzy:!0,ebonPlaguebringer:!0,faerieFire:S.TristateEffectImproved,heartOfTheCrusader:!0,judgementOfWisdom:!0,shadowMastery:!0,sunderArmor:!0,totemOfWrath:!0}),Q={distanceFromTarget:18},$={name:"P2 Preset",tooltip:B,gear:w.fromJsonString(' {\n      "items": [\n        {\n          "id": 45497,\n          "enchant": 3820,\n          "gems": [\n            41285,\n            42144\n          ]\n        },\n        {\n          "id": 45133,\n          "gems": [\n            40048\n          ]\n        },\n        {\n          "id": 46196,\n          "enchant": 3810,\n          "gems": [\n            39998\n          ]\n        },\n        {\n          "id": 45242,\n          "enchant": 3859,\n          "gems": [\n            40048\n          ]\n        },\n        {\n          "id": 45519,\n          "enchant": 3832,\n          "gems": [\n            40051,\n            42144,\n            40026\n          ]\n        },\n        {\n          "id": 45446,\n          "enchant": 2332,\n          "gems": [\n            42144,\n            0\n          ]\n        },\n        {\n          "id": 45665,\n          "enchant": 3604,\n          "gems": [\n            39998,\n            39998,\n            0\n          ]\n        },\n        {\n          "id": 45619,\n          "gems": [\n            39998,\n            39998,\n            39998\n          ]\n        },\n        {\n          "id": 46192,\n          "enchant": 3719,\n          "gems": [\n            39998,\n            39998\n          ]\n        },\n        {\n          "id": 45537,\n          "enchant": 3606,\n          "gems": [\n            39998,\n            40026\n          ]\n        },\n        {\n          "id": 46046,\n          "gems": [\n            39998\n          ]\n        },\n        {\n          "id": 45495,\n          "gems": [\n            39998\n          ]\n        },\n        {\n          "id": 45466\n        },\n        {\n          "id": 45518\n        },\n        {\n          "id": 45620,\n          "enchant": 3834,\n          "gems": [\n            39998\n          ]\n        },\n        {\n          "id": 45617\n        },\n        {\n          "id": 40321\n        }\n      ]\n    }')},X={name:"P1 Preset",tooltip:B,gear:w.fromJsonString('{"items": [\n\t\t{\n\t\t\t"id": 40467,\n\t\t\t"enchant": 3820,\n\t\t\t"gems": [\n\t\t\t\t41285,\n\t\t\t\t42144\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 44661,\n\t\t\t"gems": [\n\t\t\t\t40026\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 40470,\n\t\t\t"enchant": 3810,\n\t\t\t"gems": [\n\t\t\t\t42144\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 44005,\n\t\t\t"enchant": 3859,\n\t\t\t"gems": [\n\t\t\t\t40026\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 40469,\n\t\t\t"enchant": 3832,\n\t\t\t"gems": [\n\t\t\t\t42144,\n\t\t\t\t39998\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 44008,\n\t\t\t"enchant": 2332,\n\t\t\t"gems": [\n\t\t\t\t39998,\n\t\t\t\t0\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 40466,\n\t\t\t"enchant": 3604,\n\t\t\t"gems": [\n\t\t\t\t39998,\n\t\t\t\t0\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 40561,\n\t\t\t"enchant": 3601,\n\t\t\t"gems": [\n\t\t\t\t39998\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 40560,\n\t\t\t"enchant": 3719\n\t\t},\n\t\t{\n\t\t\t"id": 40519,\n\t\t\t"enchant": 3606\n\t\t},\n\t\t{\n\t\t\t"id": 40399\n\t\t},\n\t\t{\n\t\t\t"id": 40080\n\t\t},\n\t\t{\n\t\t\t"id": 40255\n\t\t},\n\t\t{\n\t\t\t"id": 40432\n\t\t},\n\t\t{\n\t\t\t"id": 40395,\n\t\t\t"enchant": 3834\n\t\t},\n\t\t{\n\t\t\t"id": 40192\n\t\t},\n\t\t{\n\t\t\t"id": 40321\n\t\t}\n\t]}')},Y={name:"Pre-raid Preset",tooltip:B,gear:w.fromJsonString('{ "items": [\n\t\t{\n\t\t\t"id": 42554,\n\t\t\t"enchant": 3820,\n\t\t\t"gems": [\n\t\t\t\t41285,\n\t\t\t\t40049\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 40680\n\t\t},\n\t\t{\n\t\t\t"id": 37673,\n\t\t\t"enchant": 3810,\n\t\t\t"gems": [\n\t\t\t\t42144\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 41610,\n\t\t\t"enchant": 3859\n\t\t},\n\t\t{\n\t\t\t"id": 39547,\n\t\t\t"enchant": 3832,\n\t\t\t"gems": [\n\t\t\t\t42144,\n\t\t\t\t40026\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 37884,\n\t\t\t"enchant": 2332,\n\t\t\t"gems": [\n\t\t\t\t0\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 39544,\n\t\t\t"enchant": 3604,\n\t\t\t"gems": [\n\t\t\t\t42144,\n\t\t\t\t0\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 40696,\n\t\t\t"enchant": 3601,\n\t\t\t"gems": [\n\t\t\t\t40014,\n\t\t\t\t39998\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 37854,\n\t\t\t"enchant": 3719\n\t\t},\n\t\t{\n\t\t\t"id": 44202,\n\t\t\t"enchant": 3606,\n\t\t\t"gems": [\n\t\t\t\t39998\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 40585\n\t\t},\n\t\t{\n\t\t\t"id": 43253,\n\t\t\t"gems": [\n\t\t\t\t40026\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t"id": 37873\n\t\t},\n\t\t{\n\t\t\t"id": 40682\n\t\t},\n\t\t{\n\t\t\t"id": 45085,\n\t\t\t"enchant": 3834\n\t\t},\n\t\t{\n\t\t\t"id": 40698\n\t\t},\n\t\t{\n\t\t\t"id": 40712\n\t\t}\n\t]}')};class Z extends C{constructor(t,n){super(t,n,{cssClass:"balance-druid-sim-ui",cssScheme:"druid",knownIssues:[],epStats:[v.StatIntellect,v.StatSpirit,v.StatSpellPower,v.StatSpellHit,v.StatSpellCrit,v.StatSpellHaste,v.StatMP5],epReferenceStat:v.StatSpellPower,displayStats:[v.StatHealth,v.StatStamina,v.StatIntellect,v.StatSpirit,v.StatSpellPower,v.StatSpellHit,v.StatSpellCrit,v.StatSpellHaste,v.StatMP5],defaults:{gear:X.gear,epWeights:I.fromMap({[v.StatIntellect]:.38,[v.StatSpirit]:.34,[v.StatSpellPower]:1,[v.StatSpellCrit]:.69,[v.StatSpellHaste]:.77,[v.StatMP5]:0}),consumes:J,rotation:G,talents:z.data,specOptions:L,raidBuffs:_,partyBuffs:V,individualBuffs:K,debuffs:q,other:Q},playerIconInputs:[k],rotationInputs:j,includeBuffDebuffInputs:[R,x,N,U,F,A,D],excludeBuffDebuffInputs:[],otherInputs:{inputs:[E,H]},encounterPicker:{showExecuteProportion:!1},presets:{talents:[z],gear:[Y,X,$]}})}}export{Z as B,G as D,Q as O,X as P,z as S,L as a,J as b,$ as c};
//# sourceMappingURL=sim-5bcf69ad.chunk.js.map

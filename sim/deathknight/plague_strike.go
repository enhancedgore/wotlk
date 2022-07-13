package deathknight

import (
	"github.com/wowsims/wotlk/sim/core"
)

func (deathKnight *DeathKnight) registerPlagueStrikeSpell() {
	weaponBaseDamage := core.BaseDamageFuncMeleeWeapon(core.MainHand, true, 0.0, 0.5, true)

	deathKnight.PlagueStrike = deathKnight.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 49921},
		SpellSchool: core.SpellSchoolPhysical,
		Flags:       core.SpellFlagMeleeMetrics,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask:         core.ProcMaskMeleeMHSpecial,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			BaseDamage: core.BaseDamageConfig{
				Calculator: func(sim *core.Simulation, hitEffect *core.SpellEffect, spell *core.Spell) float64 {
					return weaponBaseDamage(sim, hitEffect, spell)
				},
				TargetSpellCoefficient: 1,
			},

			OutcomeApplier: deathKnight.OutcomeFuncMeleeSpecialHitAndCrit(deathKnight.DefaultMeleeCritMultiplier()),

			OnSpellHitDealt: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					dkSpellCost := deathKnight.DetermineOptimalCost(sim, 0, 0, 1)
					deathKnight.Spend(sim, spell, dkSpellCost)

					deathKnight.BloodPlagueDisease.Apply(sim)

					deathKnight.AddRunicPower(sim, 10.0, spell.RunicPowerMetrics())
				}
			},
		}),
	})
}

func (deathKnight *DeathKnight) CanPlagueStrike(sim *core.Simulation) bool {
	return deathKnight.CastCostPossible(sim, 0.0, 0, 0, 1) && deathKnight.PlagueStrike.IsReady(sim)
}
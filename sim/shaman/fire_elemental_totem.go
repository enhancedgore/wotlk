package shaman

import (
	"time"

	"github.com/wowsims/wotlk/sim/core"
	"github.com/wowsims/wotlk/sim/core/proto"
	"github.com/wowsims/wotlk/sim/core/stats"
)

const fireTotemDuration time.Duration = time.Second * 120

func (shaman *Shaman) registerFireElementalTotem() {
	if !shaman.Totems.UseFireElemental {
		return
	}

	actionID := core.ActionID{SpellID: 2894}
	manaCost := 0.23 * shaman.BaseMana

	fireElementalAura := shaman.RegisterAura(core.Aura{
		Label:    "Fire Elemental Totem",
		ActionID: actionID,
		Duration: fireTotemDuration,
	})

	shaman.FireElementalTotem = shaman.RegisterSpell(core.SpellConfig{
		ActionID: actionID,

		ResourceType: stats.Mana,
		BaseCost:     manaCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: manaCost,
				GCD:  core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    shaman.NewTimer(),
				Duration: time.Minute * time.Duration(core.TernaryFloat64(shaman.HasMajorGlyph(proto.ShamanMajorGlyph_GlyphOfFireElementalTotem), 5, 10)),
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			// TODO: ToW needs a unique buff/debuff aura for each raidmember/target.
			//  Otherwise we will be possibly disabling another ele shaman's ToW debuff/buff.
			if shaman.Totems.Fire == proto.FireTotem_TotemOfWrath {
				shaman.NextTotemDrops[FireTotem] = sim.CurrentTime + fireTotemDuration
			} else if shaman.Totems.Fire != proto.FireTotem_NoFireTotem && !shaman.Totems.UseFireMcd {
				shaman.NextTotemDrops[FireTotem] = sim.CurrentTime + fireTotemDuration
			}
			shaman.MagmaTotemDot.Cancel(sim)
			shaman.SearingTotemDot.Cancel(sim)
			shaman.FireElemental.EnableWithTimeout(sim, shaman.FireElemental, fireTotemDuration)

			// Add a dummy aura to show in metrics
			fireElementalAura.Activate(sim)
		},
	})

	shaman.AddMajorCooldown(core.MajorCooldown{
		Spell: shaman.FireElementalTotem,
		Type:  core.CooldownTypeUnknown,
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			success := shaman.fireElementalSnapShot.CanSnapShot(sim, shaman.SpellGCD()+50*time.Millisecond)

			if (sim.Encounter.Duration <= 120*time.Second && sim.CurrentTime >= 10*time.Second) || (sim.Encounter.Duration > 120*time.Second && sim.CurrentTime >= 20*time.Second) {
				success = true
			}

			if success {
				shaman.castFireElemental = true
				shaman.fireElementalSnapShot.ActivateMajorCooldowns(sim)
				shaman.fireElementalSnapShot.ResetProcTrackers()
				shaman.castFireElemental = false
			}

			return success
		},
	})
}

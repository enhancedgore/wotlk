package dps

import (
	"time"

	"github.com/wowsims/wotlk/sim/core"
	"github.com/wowsims/wotlk/sim/deathknight"
)

var UAsoon bool = false
var UAsoonBS bool = false

func (dk *DpsDeathknight) RotationActionCallback_FrostSubBlood_TrySequence(sim *core.Simulation, target *core.Unit, s *deathknight.Sequence) time.Duration {
	UACheckPesti := false
	UACheckBS := false
	EOFCheck := false
	UAsoon = false
	UAsoonBS = false
	UACheckPesti = dk.FrostSubBlood_UACheck(sim, target, s, "Pestilence")
	UACheckBS = dk.FrostSubBlood_UACheck(sim, target, s, "Blood Strike")
	EOFCheck = dk.FrostSubBlood_EOFCheck(sim, target, s)
	bothblAt := dk.BloodDeathRuneBothReadyAt()
	ffExpiresAt := dk.FrostFeverSpell.Dot(target).ExpiresAt()
	bpExpiresAt := dk.BloodPlagueSpell.Dot(target).ExpiresAt()
	diseaseExpiresAt := core.MinDuration(ffExpiresAt, bpExpiresAt)

	if EOFCheck {
		s.Clear().NewAction(dk.RotationActionCallback_EndOfFightPrio)
	} else if UACheckPesti {
		if dk.UnbreakableArmor.IsReady(sim) && (dk.CurrentBloodRunes() >= 1 || dk.CurrentDeathRunes() >= 1) {
			//use UA now
			s.Clear().NewAction(dk.RotationActionCallback_FrostSubBlood_UA_Now)
		} else if dk.UnbreakableArmor.TimeToReady(sim)+sim.CurrentTime < diseaseExpiresAt && dk.UnbreakableArmor.TimeToReady(sim)+sim.CurrentTime > bothblAt {
			UAsoon = true
			s.Clear().NewAction(dk.RotationActionCallback_FrostSubBlood_NormalPrio)
		}
	} else {
		if UACheckBS {
			UAsoonBS = true
		}
		s.Clear().NewAction(dk.RotationActionCallback_FrostSubBlood_NormalPrio)
	}
	return sim.CurrentTime
}

// UA check
func (dk *DpsDeathknight) FrostSubBlood_UACheck(sim *core.Simulation, target *core.Unit, s *deathknight.Sequence, spell string) bool {
	bothblAt := dk.BloodDeathRuneBothReadyAt()
	ffExpiresAt := dk.FrostFeverSpell.Dot(target).ExpiresAt()
	bpExpiresAt := dk.BloodPlagueSpell.Dot(target).ExpiresAt()
	diseaseExpiresAt := core.MinDuration(ffExpiresAt, bpExpiresAt)
	if spell == "Pestilence" {
		if dk.UnbreakableArmor.IsReady(sim) && (dk.CurrentBloodRunes()+dk.CurrentDeathRunes() > 1) || dk.UnbreakableArmor.TimeToReady(sim)+sim.CurrentTime < diseaseExpiresAt && dk.UnbreakableArmor.TimeToReady(sim)+sim.CurrentTime > bothblAt {
			return true
		}
		return false
	} else if spell == "Blood Strike" {
		return dk.UnbreakableArmor.TimeToReady(sim) < 10*time.Second
	}
	return false
}

// EOF check
func (dk *DpsDeathknight) FrostSubBlood_EOFCheck(sim *core.Simulation, target *core.Unit, s *deathknight.Sequence) bool {
	if sim.CurrentTime+7000*time.Millisecond > sim.CurrentTime+sim.GetRemainingDuration() {
		return true
	} else {
		return false
	}
}

func (dk *DpsDeathknight) RotationActionCallback_FrostSubBlood_UA_Now(sim *core.Simulation, target *core.Unit, s *deathknight.Sequence) time.Duration {
	s.Clear().
		NewAction(dk.RotationActionCallback_UA_Frost).
		NewAction(dk.RotationActionCallback_Pesti).
		NewAction(dk.RotationActionCallback_FrostSubBlood_TrySequence)
	return sim.CurrentTime
}

// Normal rotation
func (dk *DpsDeathknight) RotationActionCallback_FrostSubBlood_Rotation(sim *core.Simulation, target *core.Unit, s *deathknight.Sequence) time.Duration {
	s.Clear().NewAction(dk.RotationActionCallback_FrostSubBlood_NormalPrio)
	return sim.CurrentTime
}

func (dk *DpsDeathknight) setupFrostSubBloodERWOpener_v3() {
	dk.setupUnbreakableArmorCooldowns()

	dk.RotationSequence.
		NewAction(dk.RotationActionCallback_IT).
		NewAction(dk.RotationActionCallback_PS).
		NewAction(dk.RotationActionCallback_UA_Frost).
		NewAction(dk.RotationActionCallback_BT).
		NewAction(dk.RotationActionCallback_FrostSubBlood_Obli).
		NewAction(dk.RotationActionCallback_FS).
		NewAction(dk.RotationActionCallback_FrostSubBlood_Sequence_Pesti).
		NewAction(dk.RotationActionCallback_ERW).
		NewAction(dk.RotationActionCallback_FrostSubBlood_Obli).
		NewAction(dk.RotationActionCallback_FrostSubBlood_Obli).
		NewAction(dk.RotationActionCallback_FrostSubBlood_Obli).
		NewAction(dk.RotationActionCallback_Frost_FS_HB).
		NewAction(dk.RotationActionCallback_RD).
		NewAction(dk.RotationActionCallback_Frost_FS_HB).
		NewAction(dk.RotationActionCallback_Frost_FS_HB).
		NewAction(dk.RotationActionCallback_FrostSubBlood_Obli).
		NewAction(dk.RotationActionCallback_FrostSubBlood_Obli).
		NewAction(dk.RotationActionCallback_FrostSubBlood_Sequence_Pesti).
		NewAction(dk.RotationActionCallback_FS).
		NewAction(dk.RotationActionCallback_BS).
		NewAction(dk.RotationActionCallback_FrostSubBlood_TrySequence)
}

func (dk *DpsDeathknight) FrostSubBlood_runeTimeCheck(sim *core.Simulation, target *core.Unit, s *deathknight.Sequence) bool {
	frAt := dk.NormalFrostRuneReadyAt(sim)
	uhAt := dk.NormalUnholyRuneReadyAt(sim)
	obAt := core.MaxDuration(frAt, uhAt)
	abGcd := 1500 * time.Millisecond
	spGcd := dk.SpellGCD()
	delayAmount := 2501 * time.Millisecond
	casted := false

	if obAt+delayAmount > sim.CurrentTime+spGcd {
		if dk.RimeAura.IsActive() {
			casted = true
			dk.HowlingBlast.Cast(sim, target)
		} else if dk.HornOfWinter.CanCast(sim, nil) && dk.CurrentRunicPower() < dk.MaxRunicPower()-2.0*dk.fr.oblitRPRegen-10 {
			casted = true
			dk.HornOfWinter.Cast(sim, target)
		}
	}
	if !casted {
		if obAt+delayAmount > sim.CurrentTime+abGcd {
			if dk.FrostStrike.CanCast(sim, nil) {
				casted = true
				dk.FrostStrike.Cast(sim, target)
			}
		}
	}
	s.Clear().NewAction(dk.RotationActionCallback_FrostSubBlood_TrySequence)
	return casted
}

func (dk *DpsDeathknight) RotationActionCallback_FrostSubBlood_NormalPrio(sim *core.Simulation, target *core.Unit, s *deathknight.Sequence) time.Duration {
	ffExpiresAt := dk.FrostFeverSpell.Dot(target).ExpiresAt()
	bpExpiresAt := dk.BloodPlagueSpell.Dot(target).ExpiresAt()
	diseaseExpiresAt := core.MinDuration(ffExpiresAt, bpExpiresAt)
	bothblAt := dk.BloodDeathRuneBothReadyAt()

	frAt := dk.NormalFrostRuneReadyAt(sim)
	uhAt := dk.NormalUnholyRuneReadyAt(sim)
	obAt := core.MaxDuration(frAt, uhAt)
	abGcd := 1500 * time.Millisecond
	delayAmount := 2501 * time.Millisecond

	dk.RunicPowerBar.CopyRunicPowerBar()
	if ffExpiresAt < sim.CurrentTime+dk.SpellGCD() {
		dk.IcyTouch.Cast(sim, target)
		s.Clear().NewAction(dk.RotationActionCallback_FrostSubBlood_TrySequence)
	} else if bpExpiresAt < sim.CurrentTime+abGcd {
		dk.PlagueStrike.Cast(sim, target)
		s.Clear().NewAction(dk.RotationActionCallback_FrostSubBlood_TrySequence)
	} else if diseaseExpiresAt < sim.CurrentTime+5000*time.Millisecond && dk.Pestilence.CanCast(sim, nil) && !UAsoon { //no rune grace yet
		dk.Pestilence.Cast(sim, target)
		s.Clear().NewAction(dk.RotationActionCallback_FrostSubBlood_TrySequence)
		return sim.CurrentTime
	} else if dk.Obliterate.CanCast(sim, nil) && (dk.CurrentFrostRunes() >= 1 && dk.CurrentUnholyRunes() >= 1) {
		dk.Obliterate.Cast(sim, target)
		s.Clear().NewAction(dk.RotationActionCallback_FrostSubBlood_TrySequence)
		return sim.CurrentTime
	} else if dk.Obliterate.CanCast(sim, nil) && dk.UnbreakableArmor.CD.TimeToReady(sim) > 35*time.Second {
		dk.Obliterate.Cast(sim, target)
		if dk.BloodTap.CanCast(sim, nil) {
			dk.BloodTap.Cast(sim, target)
		}
		s.Clear().NewAction(dk.RotationActionCallback_FrostSubBlood_TrySequence)
		return sim.CurrentTime
	} else if dk.BloodStrike.CanCast(sim, nil) && dk.CurrentRunicPower() < 70 && (dk.CurrentBloodRunes()+dk.CurrentDeathRunes() > 1 || diseaseExpiresAt > 10*time.Second+sim.CurrentTime) && (sim.CurrentTime+abGcd < obAt+delayAmount) && !UAsoonBS { //70 is arbitrary for now, it should be changed to max - 2 oblit - 1 bs
		dk.BloodStrike.Cast(sim, target)
		s.Clear().NewAction(dk.RotationActionCallback_FrostSubBlood_TrySequence)
		return sim.CurrentTime
	} else if dk.BloodStrike.CanCast(sim, nil) && dk.UnbreakableArmorAura.IsActive() && dk.UnbreakableArmorAura.RemainingDuration(sim)+sim.CurrentTime > 10*time.Second+bothblAt {
		dk.BloodStrike.Cast(sim, target)
		s.Clear().
			NewAction(dk.RotationActionCallback_Pesti).
			NewAction(dk.RotationActionCallback_FrostSubBlood_TrySequence)
		return sim.CurrentTime
	} else if dk.FrostSubBlood_runeTimeCheck(sim, target, s) {
		return sim.CurrentTime
	}
	s.Clear().NewAction(dk.RotationActionCallback_FrostSubBlood_TrySequence)
	return -1
}

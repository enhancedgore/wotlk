package dps

import (
	"time"

	"github.com/wowsims/wotlk/sim/core"
	"github.com/wowsims/wotlk/sim/deathknight"
)

func (dk *DpsDeathknight) RotationActionCallback_FrostSubBlood_TrySequence(sim *core.Simulation, target *core.Unit, s *deathknight.Sequence) time.Duration {
	UACheck := false
	EOFCheck := false
	UACheck = dk.FrostSubBlood_UACheck(sim, target, s)
	EOFCheck = dk.FrostSubBlood_EOFCheck(sim, target, s)

	if UACheck {
		if dk.UnbreakableArmor.IsReady(sim) && (dk.CurrentBloodRunes() >= 1 || dk.CurrentDeathRunes() >= 1) {
			//use UA now
			s.Clear().NewAction(dk.RotationActionCallback_FrostSubBlood_UA_Now)
		} else if dk.UnbreakableArmor.TimeToReady(sim) <= 10*time.Second {
			//prep UA now, set a new variable for UA prep and do not allow pesti
			s.Clear().NewAction(dk.RotationActionCallback_FrostSubBlood_NormalPrio)
		}
	} else if EOFCheck {
		s.Clear().NewAction(dk.RotationActionCallback_EndOfFightPrio)
	} else {
		s.Clear().NewAction(dk.RotationActionCallback_FrostSubBlood_NormalPrio)
	}
	return sim.CurrentTime
}

// UA check
func (dk *DpsDeathknight) FrostSubBlood_UACheck(sim *core.Simulation, target *core.Unit, s *deathknight.Sequence) bool {
	if (dk.UnbreakableArmor.IsReady(sim) && (dk.CurrentBloodRunes()+dk.CurrentDeathRunes() > 1)) || dk.UnbreakableArmor.TimeToReady(sim) <= 10*time.Second {
		return true
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

func (dk *DpsDeathknight) setupFrostSubBloodERWOpener_v2() {
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
	delayAmount := core.MinDuration(time.Duration(dk.Rotation.OblitDelayDuration)*time.Millisecond, 2501*time.Millisecond)
	casted := false

	if obAt+delayAmount > sim.CurrentTime+spGcd {
		if dk.Rime() {
			casted = true
			dk.HowlingBlast.Cast(sim, target)
		} else if dk.HornOfWinter.CanCast(sim) && dk.CurrentRunicPower() < dk.MaxRunicPower()-2.0*dk.fr.oblitRPRegen-10 {
			casted = true
			dk.HornOfWinter.Cast(sim, target)
		}
	}
	if !casted {
		if obAt+delayAmount > sim.CurrentTime+abGcd {
			if dk.FrostStrike.CanCast(sim) {
				casted = true
				dk.FrostStrike.Cast(sim, target)
			}
		}
	}
	s.Clear().NewAction(dk.RotationActionCallback_FrostSubBlood_TrySequence)
	return casted
}

func (dk *DpsDeathknight) RotationActionCallback_FrostSubBlood_NormalPrio(sim *core.Simulation, target *core.Unit, s *deathknight.Sequence) time.Duration {
	ffExpiresAt := dk.FrostFeverDisease[target.Index].ExpiresAt()
	bpExpiresAt := dk.BloodPlagueDisease[target.Index].ExpiresAt()
	diseaseExpiresAt := core.MinDuration(ffExpiresAt, bpExpiresAt)

	frAt := dk.NormalFrostRuneReadyAt(sim)
	uhAt := dk.NormalUnholyRuneReadyAt(sim)
	obAt := core.MaxDuration(frAt, uhAt)
	abGcd := 1500 * time.Millisecond
	delayAmount := core.MinDuration(time.Duration(dk.Rotation.OblitDelayDuration)*time.Millisecond, 2501*time.Millisecond)

	dk.RunicPowerBar.CopyRunicPowerBar()

	if diseaseExpiresAt < sim.CurrentTime+4*time.Second && dk.Pestilence.CanCast(sim) { //no rune grace yet
		dk.Pestilence.Cast(sim, target)
		s.Clear().NewAction(dk.RotationActionCallback_FrostSubBlood_TrySequence)
		return sim.CurrentTime
	} else if dk.Obliterate.CanCast(sim) && dk.CurrentFrostRunes() >= 1 && dk.CurrentUnholyRunes() >= 1 {
		dk.Obliterate.Cast(sim, target)
		s.Clear().NewAction(dk.RotationActionCallback_FrostSubBlood_TrySequence)
		return sim.CurrentTime
	} else if dk.BloodStrike.CanCast(sim) && dk.CurrentRunicPower() < 70 && (dk.CurrentBloodRunes()+dk.CurrentDeathRunes() > 1 || diseaseExpiresAt > 10*time.Second+sim.CurrentTime) && (sim.CurrentTime+abGcd < obAt+delayAmount) { //70 is arbitrary for now, it should be changed to max - 2 oblit - 1 bs
		dk.BloodStrike.Cast(sim, target)
		s.Clear().NewAction(dk.RotationActionCallback_FrostSubBlood_TrySequence)
		return sim.CurrentTime
	} else if dk.FrostSubBlood_runeTimeCheck(sim, target, s) {
		return sim.CurrentTime
	}
	s.Clear().NewAction(dk.RotationActionCallback_FrostSubBlood_TrySequence)
	return -1
}

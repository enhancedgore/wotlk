package dps

import (
	"time"

	"github.com/wowsims/wotlk/sim/core"
	"github.com/wowsims/wotlk/sim/deathknight"
)

// end of fight oblit does not check diseases, it just presses it regardless, but will retry if fails to land.
func (dk *DpsDeathknight) RotationActionCallback_FrostSubUnh_EndOfFight_Obli(sim *core.Simulation, target *core.Unit, s *deathknight.Sequence) time.Duration {
	casted := false
	advance := true
	waitTime := time.Duration(-1)
	if dk.Obliterate.CanCast(sim) {
		if dk.Deathchill != nil && dk.Deathchill.IsReady(sim) {
			dk.Deathchill.Cast(sim, target)
		}
		casted = dk.Obliterate.Cast(sim, target)
		advance = dk.LastOutcome.Matches(core.OutcomeLanded)
	}
	s.ConditionalAdvance(casted && advance)
	return core.TernaryDuration(casted, -1, waitTime)
}

func (dk *DpsDeathknight) RegularPrioPickSpell(sim *core.Simulation, target *core.Unit, untilTime time.Duration) *deathknight.RuneSpell {
	fsCost := float64(core.RuneCost(dk.FrostStrike.CurCast.Cost).RunicPower())

	abGcd := 1500 * time.Millisecond
	spGcd := dk.SpellGCD()

	km := dk.KM()
	rime := dk.Rime()
	if sim.CurrentTime+abGcd <= untilTime && dk.FrostStrike.CanCast(sim) && km {
		return dk.FrostStrike
	} else if sim.CurrentTime+abGcd <= untilTime && dk.FrostStrike.CanCast(sim) && dk.CurrentRunicPower() >= 100.0 {
		return dk.FrostStrike
	} else if sim.CurrentTime+spGcd <= untilTime && dk.HowlingBlast.CanCast(sim) && rime {
		return dk.HowlingBlast
	} else if sim.CurrentTime+abGcd <= untilTime && dk.FrostStrike.CanCast(sim) && dk.CurrentRunicPower() >= 2.0*(fsCost-dk.fr.oblitRPRegen) {
		return dk.FrostStrike
	} else if sim.CurrentTime+spGcd <= untilTime && dk.HornOfWinter.CanCast(sim) {
		return dk.HornOfWinter
	} else {
		return nil
	}
}

//end of fight functions coming soon

func (dk *DpsDeathknight) RotationActionCallback_EndOfFightCheck(sim *core.Simulation, target *core.Unit, s *deathknight.Sequence) time.Duration {
	if sim.CurrentTime /*+10000*time.Millisecond*/ > sim.GetMaxDuration() {
		s.Clear().NewAction(dk.RotationActionCallback_EndOfFightPrio)
	} else {
		s.Advance()
	}
	return sim.CurrentTime
}

func (dk *DpsDeathknight) RotationActionCallback_EndOfFightPrio(sim *core.Simulation, target *core.Unit, s *deathknight.Sequence) time.Duration {

	ffExpiresAt := dk.FrostFeverDisease[target.Index].ExpiresAt()
	bpExpiresAt := dk.BloodPlagueDisease[target.Index].ExpiresAt()
	diseaseExpiresAt := core.MinDuration(ffExpiresAt, bpExpiresAt)
	abGcd := 1500 * time.Millisecond
	spGcd := dk.SpellGCD()
	frAt := dk.NormalFrostRuneReadyAt(sim)
	uhAt := dk.NormalUnholyRuneReadyAt(sim)
	obAt := core.MaxDuration(frAt, uhAt)
	fsCost := float64(core.RuneCost(dk.FrostStrike.CurCast.Cost).RunicPower())
	delayAmount := core.MinDuration(time.Duration(dk.Rotation.OblitDelayDuration)*time.Millisecond, 2501*time.Millisecond)
	//call dk.CurrentRunicPower()
	//call rp.runeMeta

	//diseases last until end of fight
	if diseaseExpiresAt >= sim.GetMaxDuration() {
		//if can cast oblit
		if sim.CurrentTime >= obAt {
			s.Clear().
				NewAction(dk.RotationActionCallback_FrostSubUnh_EndOfFight_Obli).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if sim.CurrentTime+spGcd > sim.GetMaxDuration() && obAt < sim.GetMaxDuration() {
			//wait to cast oblit
			s.Clear().
				NewAction(dk.RotationActionCallback_FrostSubUnh_EndOfFight_Obli).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if sim.CurrentTime+spGcd < sim.GetMaxDuration() && sim.CurrentTime+abGcd > sim.GetMaxDuration() && obAt < sim.GetMaxDuration() && dk.Rime() {
			//if you can only cast a spell GCD to catch the last oblit before fight ends, and have rime
			//cast rime
			s.Clear().
				NewAction(dk.RotationActionCallback_HB).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if (sim.CurrentTime+abGcd > sim.GetMaxDuration() || sim.CurrentTime+abGcd > obAt+delayAmount) && obAt < sim.GetMaxDuration() {
			//wait to cast oblit
			s.Clear().
				NewAction(dk.RotationActionCallback_FrostSubUnh_EndOfFight_Obli).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if dk.CurrentRunicPower() >= fsCost && sim.CurrentTime+abGcd < obAt+delayAmount {
			s.Clear().
				NewAction(dk.RotationActionCallback_FS).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if dk.Rime() && sim.CurrentTime+spGcd < obAt+delayAmount {
			//cast HB
			s.Clear().
				NewAction(dk.RotationActionCallback_HB).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if (dk.CurrentBloodRunes() >= 1 || dk.CurrentDeathRunes() == 1) && (sim.CurrentTime+abGcd < obAt+delayAmount) {
			s.Clear().
				NewAction(dk.RotationActionCallback_BS).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if dk.HornOfWinter.IsReady(sim) && sim.CurrentTime+spGcd < obAt+delayAmount {
			//cast horn
			s.Clear().
				NewAction(dk.RotationActionCallback_HW).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else {
			s.Clear().
				NewAction(dk.RotationActionCallback_FrostSubUnh_EndOfFight_Obli).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		}
	} else if diseaseExpiresAt >= sim.GetMaxDuration()-abGcd { //disease expires less than 1 gcd before end of fight
		//if can cast oblit
		if sim.CurrentTime >= obAt {
			//cast oblit
			s.Clear().
				NewAction(dk.RotationActionCallback_FrostSubUnh_EndOfFight_Obli).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if sim.CurrentTime+spGcd > sim.GetMaxDuration() && obAt < sim.GetMaxDuration() {
			//wait to cast oblit
			s.Clear().
				NewAction(dk.RotationActionCallback_FrostSubUnh_EndOfFight_Obli).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if sim.CurrentTime+spGcd < sim.GetMaxDuration() && sim.CurrentTime+abGcd > sim.GetMaxDuration() && obAt < sim.GetMaxDuration() && dk.Rime() {
			//if you can only cast a spell GCD to catch the last oblit before fight ends, and have rime
			//cast rime
			s.Clear().
				NewAction(dk.RotationActionCallback_HB).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if sim.CurrentTime+abGcd > sim.GetMaxDuration() && obAt < sim.GetMaxDuration() {
			//wait to cast oblit
			s.Clear().
				NewAction(dk.RotationActionCallback_FrostSubUnh_EndOfFight_Obli).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if dk.CurrentRunicPower() >= fsCost {
			//cast FS
			s.Clear().
				NewAction(dk.RotationActionCallback_FS).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if dk.Rime() {
			//cast HB
			s.Clear().
				NewAction(dk.RotationActionCallback_HB).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if dk.CurrentBloodRunes() >= 1 || sim.CurrentTime < diseaseExpiresAt {
			//cast pesti
			s.Clear().
				NewAction(dk.RotationActionCallback_Pesti).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if dk.CurrentBloodRunes() >= 1 {
			//cast BS
			s.Clear().
				NewAction(dk.RotationActionCallback_BS).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if dk.HornOfWinter.IsReady(sim) {
			//cast horn
			s.Clear().
				NewAction(dk.RotationActionCallback_HW).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		}
	} else if sim.CurrentTime+spGcd < diseaseExpiresAt && sim.CurrentTime+abGcd > diseaseExpiresAt && dk.CurrentRunicPower() < 100 && (dk.Rime() || dk.HornOfWinter.IsReady(sim)) { //if you can fit a spellgcd before disease dropping
		if dk.Rime() && dk.CurrentRunicPower() < 100 {
			s.Clear().
				NewAction(dk.RotationActionCallback_HB).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if dk.HornOfWinter.IsReady(sim) && dk.CurrentRunicPower() < 100 {
			s.Clear().
				NewAction(dk.RotationActionCallback_HW).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		}
	} else if sim.CurrentTime+abGcd < diseaseExpiresAt {
		//if can cast oblit
		if sim.CurrentTime >= obAt {
			//cast oblit
			s.Clear().
				NewAction(dk.RotationActionCallback_FrostSubUnh_EndOfFight_Obli).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if sim.CurrentTime+spGcd > obAt+delayAmount {
			//wait to cast oblit
			s.Clear().
				NewAction(dk.RotationActionCallback_FrostSubUnh_EndOfFight_Obli).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if sim.CurrentTime+spGcd < obAt+delayAmount && sim.CurrentTime+abGcd > obAt && (dk.Rime() || dk.CurrentRunicPower() < fsCost*4-2*dk.fr.oblitRPRegen && dk.HornOfWinter.IsReady(sim)) {
			if dk.Rime() {
				//cast HB
				s.Clear().
					NewAction(dk.RotationActionCallback_HB).
					NewAction(dk.RotationActionCallback_EndOfFightCheck)
			} else if dk.CurrentRunicPower() < fsCost*4-2*dk.fr.oblitRPRegen && dk.HornOfWinter.IsReady(sim) { //if u wont overflow RP
				//cast horn
				s.Clear().
					NewAction(dk.RotationActionCallback_HW).
					NewAction(dk.RotationActionCallback_EndOfFightCheck)
			}
		} else if sim.CurrentTime+abGcd > obAt+delayAmount {
			//wait to cast oblit
			s.Clear().
				NewAction(dk.RotationActionCallback_FrostSubUnh_EndOfFight_Obli).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if dk.CurrentRunicPower() >= fsCost {
			//cast FS
			s.Clear().
				NewAction(dk.RotationActionCallback_FS).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if dk.Rime() {
			//cast HB
			s.Clear().
				NewAction(dk.RotationActionCallback_HB).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if (dk.CurrentBloodRunes() >= 1 || dk.CurrentDeathRunes() == 1) && diseaseExpiresAt > sim.GetMaxDuration()-abGcd {
			//cast BS
			s.Clear().
				NewAction(dk.RotationActionCallback_BS).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if (dk.CurrentBloodRunes() >= 1 || dk.CurrentDeathRunes() == 1) && diseaseExpiresAt < sim.GetMaxDuration()-abGcd {
			s.Clear().
				NewAction(dk.RotationActionCallback_Pesti).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else if dk.HornOfWinter.IsReady(sim) {
			//cast horn
			s.Clear().
				NewAction(dk.RotationActionCallback_HW).
				NewAction(dk.RotationActionCallback_EndOfFightCheck)
		} else {
			dk.WaitUntil(sim, obAt)
			s.NewAction(dk.RotationActionCallback_EndOfFightCheck)
		}
	} else if sim.CurrentTime+abGcd > diseaseExpiresAt {
		s.Clear().
			NewAction(dk.RotationActionCallback_Pesti).
			NewAction(dk.RotationActionCallback_EndOfFightCheck)
	} else {
		return -1
	}

	return sim.CurrentTime
}

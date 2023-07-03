package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGen = rand.New(randSource)
var playerHealth int = PLAYER_HEALTH
var monsterHealth int = MONSTER_HEALTH

func AttackMonster(isSpecialAttack bool) int {

	attackPower := 0
	if isSpecialAttack {
		attackPower = randomNumberGenerator(PLAYER_SPECIAL_ATTACK_MIN_DMG, PLAYER_SPECIAL_ATTACK_MAX_DMG)
	} else {
		attackPower = randomNumberGenerator(PLAYER_ATTACK_MIN_DMG, PLAYER_ATTACK_MAX_DMG)
	}
	monsterHealth -= attackPower
	// result := checkWinner()
	return attackPower
}

func AttackPlayer() int {
	attackPower := randomNumberGenerator(MONSTER_ATTACK_MIN_DMG, MONSTER_ATTACK_MAX_DMG)
	if (playerHealth - attackPower) <= 0 {
		playerHealth = 0
	} else {
		playerHealth -= attackPower
	}
	return attackPower
}
func HealPlayer() int {
	healPower := randomNumberGenerator(PLAYER_MIN_HEAL_VALUE, PLAYER_MAX_HEAL_VALUE)
	if (playerHealth + healPower) >= 100 {
		playerHealth = PLAYER_HEALTH
	} else {
		playerHealth += healPower
	}
	return healPower
}

func CheckWinner() string {
	if monsterHealth == 0 && playerHealth == 0 {
		return "DRAW"
	} else if monsterHealth > 0 && playerHealth <= 0 {
		return "MONSTER_WON"
	} else if playerHealth > 0 && monsterHealth <= 0 {
		return "PLAYER_WON "
	} else {
		return ""
	}
}

func randomNumberGenerator(min int, max int) int {
	return randGen.Intn(max-min) + min
}

func GetPlayerAndMonsterHealth() (int, int) {
	return playerHealth, monsterHealth
}

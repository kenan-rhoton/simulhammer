import random
import re

def roll(size = 6):
    return random.randrange(0,size) + 1

def title(string):
    split = re.split("[_ ]", string)
    capitalized = [w.capitalize() for w in split]
    return " ".join(capitalized)

def parse_expression(string):
    if isinstance(string, int):
        return string
    instructions = string.split()
    res = 0
    for instruction in instructions:
        if instruction == "d3":
            res += roll(3)
        elif instruction == "d6":
            res += roll()
        elif instruction.isdigit():
            res += int(instruction)
    return res


class Model:
    pass

class AttackChain:

    def __init__(self, weapon, opponent):
        self.weapon = weapon
        self.opponent = opponent
        self.hit_success = self.wound_success = False

    def execute_special_hit(self, special):
        for key in special:
            if re.fullmatch("[0-9-]", key):
                if re.fullmatch("[" + key + "]", str(self.hit_roll)):
                    self.execute_special_hit(special[key])
            if key == "mortal_wounds":
                self.opponent.damage(parse_expression(special[key]))

    def hit(self):
        if "hit" in self.weapon:
            self.hit_roll = roll()
            if self.hit_roll >= self.weapon["hit"]:
                self.hit_success = True
                if "onhit" in self.weapon:
                    onhit = self.weapon["onhit"]
                    self.execute_special_hit(self.weapon["onhit"])

    def wound(self):
        if "wound" in self.weapon and self.hit_success:
            self.wound_roll = roll()
            if self.wound_roll >= self.weapon["wound"]:
                self.wound_success = True

    def save(self):
        if self.wound_success:
            self.save_roll = roll()
            if self.save_roll < self.opponent.data["save"] - self.weapon.get("rend", 0):
                self.opponent.damage(self.weapon["damage"])

class Unit:
    def __init__(self, data, profiles):
        self.data = data
        self.models = []
        for profile in profiles:
            model = Model()
            model.leader = "leader" in profile
            if model.leader:
                profile.remove("leader")
            model.profiles = profile
            model.wounds = data["wounds"]
            self.models.append(model)

    def status(self):
        print(self.data["name"].upper())
        for model in self.models:
            stat = title(" - ".join(model.profiles)) + ": " + str(model.wounds) + " wounds"
            if model.leader:
                stat = "Leader: " + stat
            print(stat)
        print()

    def weapon_attack(self, weapon, opponent):
        for _ in range(weapon["attacks"]):
            attack_chain = AttackChain(weapon, opponent)
            attack_chain.hit()
            attack_chain.wound()
            attack_chain.save()

    def attack(self, opponent):
        for model in self.models:
            for profile in model.profiles:
                profile_data = self.data["profiles"][profile]
                self.weapon_attack(profile_data["weapon"], opponent)

    def damage(self, amount):
        while amount > 0 and len(self.models) > 0:
            model = self.models[-1]
            model.wounds -= amount
            if model.wounds <= 0:
                amount = -model.wounds
                del self.models[-1]
            else:
                amount = 0

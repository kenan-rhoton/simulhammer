import random

def roll(size = 6):
    return random.randrange(0,size) + 1

def title(string):
    split = string.split("_")
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

class Unit:
    def __init__(self, data, profiles):
        self.data = data
        self.models = []
        for profile in profiles:
            model = Model()
            model.leader = "leader" in profile
            model.profile = profile[0]
            model.wounds = data["wounds"]
            self.models.append(model)

    def status(self):
        for model in self.models:
            stat = title(model.profile) + ": " + str(model.wounds) + " wounds"
            if model.leader:
                stat = "Leader: " + stat
            print(stat)
        print()

    def weapon_attack(self, weapon, opponent):
        for _ in range(weapon["attacks"]):
            hit = roll()
            if hit >= weapon["hit"]:
                wound = roll()
                if "onhit" in weapon:
                    onhit = weapon["onhit"]
                    if "mortal_wounds" in onhit:
                        mw = parse_expression(onhit["mortal_wounds"])
                        opponent.damage(mw)
                    if "if" in onhit:
                        if hit == onhit["if"]["roll"]:
                            action = onhit["then"]
                            if "mortal_wounds" in action:
                                mw = parse_expression(action["mortal_wounds"])
                                opponent.damage(mw)
                        elif "else" in onhit:
                            action = onhit["else"]
                            if "mortal_wounds" in action:
                                mw = parse_expression(action["mortal_wounds"])
                                opponent.damage(mw)

                if "wound" in weapon and wound >= weapon["wound"]:
                    save = roll()
                    if save < opponent.data["save"] - weapon["rend"]:
                        opponent.damage(weapon["damage"])

    def attack(self, opponent):
        for model in self.models:
            profile = self.data["profiles"][model.profile]
            self.weapon_attack(profile["weapon"], opponent)
            if "base" in self.data:
                base = self.data["base"]
                if "weapon" in base:
                    self.weapon_attack(base["weapon"], opponent)

    def damage(self, amount):
        while amount > 0 and len(self.models) > 0:
            model = self.models[-1]
            model.wounds -= amount
            if model.wounds <= 0:
                amount = -model.wounds
                del self.models[-1]
            else:
                amount = 0

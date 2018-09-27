from yaml import load, dump
try:
    from yaml import CLoader as Loader, CDumper as Dumper
except:
    from yaml import Loader, Dumper

from unit import Unit

with open("data/retributors.yaml", "r") as retributors_file:
    with open("data/evocators.yaml", "r") as evocators_file:
        retributors_data = load(retributors_file, Loader=Loader)
        evocators_data = load(evocators_file, Loader=Loader)
        retributors = Unit(retributors_data,[
            ["lightning_hammer", "leader"],
            ["starsoul_mace"],
            ["starsoul_mace"],
            ["lightning_hammer"],
            ["lightning_hammer"]])
        evocators = Unit(evocators_data,[
            ["grandstave", "leader"],
            ["grandstave"],
            ["grandstave"],
            ["grandstave"],
            ["grandstave"]])
        
        print("BEFORE THE ATTACK")

        retributors.status()
        evocators.status()

        retributors.attack(evocators)

        print("AFTER THE ATTACK")

        retributors.status()
        evocators.status()

        evocators.attack(retributors)

        print("AFTER THE COUNTERATTACK")

        retributors.status()
        evocators.status()

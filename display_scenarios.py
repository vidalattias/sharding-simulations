import glob
import os

scenario_files = glob.glob('data/scenario_*.dot')

print(scenario_files)

for s in scenario_files:
    s = s.replace('data/', '').replace('.dot', '')
    os.system('dot data/' + s + '.dot -Tpdf> figures/' + s + '.pdf')
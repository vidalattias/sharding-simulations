rm data/*
rm figures/*
cd src
go build
cd ..
src/sharding-simulations

#python analysis.py
#python analyse_messages.py
#dot data/scenario_0.dot -Tpdf> figures/scenario_0.pdf
#dot data/scenario_1.dot -Tpdf> figures/scenario_1.pdf
#python display_scenarios.py

python test.py
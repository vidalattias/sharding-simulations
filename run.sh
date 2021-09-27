cd src
go build
cd ..
src/sharding-simulations
python analysis.py
dot scenario_0.dot -Tpdf> figures/scenario_0.pdf
dot scenario_1.dot -Tpdf> figures/scenario_1.pdf
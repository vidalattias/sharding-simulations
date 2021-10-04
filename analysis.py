import matplotlib.pyplot as plt

file = open("data/ratio.txt")

data = []

for l in file:
    data.append(100*float(l))


plt.ylim([0,110])
plt.plot(data)
plt.title("Ratio of messages/(messages+proofs)")
plt.savefig('figures/ratio.pdf')
plt.show()


file = open("data/messages.txt")
data = []
for l in file:
    data.append(int(l))

file = open("data/total_throughput.txt")
data_total = []
for l in file:
    data_total.append(int(l))


plt.plot(data, label="Messages only")
plt.plot(data_total, label="Total throughput")
plt.legend()
plt.title("Messages per second")
plt.savefig('figures/mps.pdf')
plt.show()
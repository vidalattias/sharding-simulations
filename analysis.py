import matplotlib.pyplot as plt

file = open("ratio.txt")

data = []

for l in file:
    data.append(100*float(l))


plt.ylim([0,110])
plt.plot(data)
plt.title("Ratio of messages/(messages+proofs)")
plt.show()


file = open("messages.txt")
data = []
for l in file:
    data.append(int(l))

file = open("total_throughput.txt")
data_total = []
for l in file:
    data_total.append(int(l))


plt.plot(data, label="Messages only")
plt.plot(data_total, label="Total throughput")
plt.legend()
plt.title("Messages per second")
plt.show()
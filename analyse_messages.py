import matplotlib.pyplot as plt
import numpy as np

f = open("txs.txt")

depths = {}
depths_time = {}

for l in f:
    arr = l.replace("\n","").split(";")
    t = float(arr[0])
    v = float(arr[1])
    depth = int(arr[2])
    proof = arr[3]

    if depth not in depths:
        depths[depth] = []
    if depth not in depths_time:
        depths_time[depth] = []

    depths[depth].append(v)
    depths_time[depth].append((t,v))



values = []
pos = []
for i in sorted(depths):
    print(np.mean(depths[i]))
    pos.append(i)
    values.append(depths[i])


plt.boxplot(values, positions = pos)

plt.ylabel("Validation time of messages (unit of time)")
plt.xlabel("Depth of tangle")

plt.show()




for i in depths_time:
    depths_time[i] = sorted(depths_time[i])
    x = []
    y = []
    for (a, b) in depths_time[i]:
        x.append(a)
        y.append(b)

    plt.plot(x, y, label="Depth "+str(i))
plt.legend()

plt.show()

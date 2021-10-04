import matplotlib.pyplot as plt
import numpy as np


step = 5

def ts(t):
    return int(t/(2*step))

def compute(i):

    f = open("data/txs_"+str(i)+".txt")

    depths = {}
    depths_time = {}

    messages = 0
    stamps = 0


    for l in f:
        arr = l.replace("\n","").split(";")
        t = float(arr[1])
        if arr[2] == "NULL":
            v = "NULL"
        else:
            v = float(arr[2])
        depth = int(arr[3])
        proof = arr[4]


        if proof == "true":
            stamps += 1
        else:
            messages += 1

        if depth not in depths:
            depths[depth] = {}
        if ts(t) not in depths[depth]:
            depths[depth][ts(t)] = [0, 0]
        
        if proof == "true":
            stamps += 1
            depths[depth][ts(t)][1] += 1
        else:
            messages += 1
            depths[depth][ts(t)][0] += 1


    total = messages + stamps


    for depth in depths:
        x = sorted(depths[depth].keys())
        y = [100*depths[depth][i][0]/(depths[depth][i][1]+depths[depth][i][0]) for i in x]
        plt.plot(x, y, label="Depth = "+str(depth))
    plt.legend()
    plt.show()



    for depth in depths:
        x = sorted(depths[depth].keys())
        y = [depths[depth][i][0] for i in x]
        plt.plot(x, y, label="Depth = "+str(depth))
    plt.legend()
    plt.show()

    print(messages)
    print(stamps)
    print(total)


    print(stamps/total*100)


compute(0)
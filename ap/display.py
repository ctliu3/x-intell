import matplotlib.pyplot as plt

centers = [13, 78, 44, 27, 50, 11]

n = 0
x = []
y = []
centerx = []
centery = []
with open('./data/testSet.txt') as f:
  for line in f:
    a = map(float, line.split())
    x.append(a[0])
    y.append(a[1])
    if n in centers:
      centerx.append(a[0])
      centery.append(a[1])
    n += 1

plt.plot(x, y, 'ro')
plt.plot(centerx, centery, 'b<')
plt.axis([min(x) - 1, max(x) + 1, min(x) - 1, max(y) + 1])
plt.show()

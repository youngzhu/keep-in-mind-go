package coin

/*
1美元（100美分）有多少种支付方式？
已知的硬币面额有：1美分，5美分，10美分，25美分，50美分

A(n)表示用1美分支付n美分
B(n)表示用1美分和5美分支付n美分
C(n)表示用1美分和5美分和10美分支付n美分
D(n)表示用1美分和5美分和10美分和25美分支付n美分
E(n)表示用1美分和5美分和10美分和25美分和50美分支付n美分

对于总额n有两种支付方式：
一，不使用50美分。支付方式有D(n)种
二，使用50美分。支付方式有E(n-50)种
所以，E(n)=D(n)+E(n-50)
同理可得，
D(n)=C(n)+D(n-25)
C(n)=B(n)+C(n-10)
B(n)=A(n)+B(n-5)
A(n)=1

令
A(0)=B(0)=C(0)=D(0)=E(0)=1
A(m)=B(m)=C(m)=D(m)=E(m)=0，m<0
*/

type Coin int

var coins = [...]Coin{1, 5, 10, 25, 50}

type change func() int

func doChange(n, i int) []change {
	var changes []change

	if n < 0 || i < 0 {
		return append(changes, negative())
	} else if n == 0 {
		return append(changes, zero())
	} else {
		f1 := doChange(n, i-1)
		changes = append(changes, f1...)

		left := n - int(coins[i])
		if left > 0 {
			f2 := doChange(n-int(coins[i]), i)
			changes = append(changes, f2...)
		}
	}
	return changes
}

func negative() change {
	return func() int {
		return 0
	}
}

func zero() change {
	return func() int {
		return 1
	}
}

func Changes(n int) int {
	changes := doChange(n, len(coins)-1)

	count := 0
	for _, change := range changes {
		count += change()
	}
	return count
}

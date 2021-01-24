package main

import "fmt"

/**
多态通过接口实现的
继承无法实现多态效果
类型转化
		//第二种断言方式
		//if rider,ok := s.(*Rider);ok {
		//	fmt.Println(rider.Attack())
		//}
*/
func main() {
	//定义一个战士的集合，加入不同的子类实现
	soldiers := make([]Soldier, 0)
	soldiers = append(soldiers, new(Rider))
	soldiers = append(soldiers, new(Master))

	//强调共性时
	fmt.Println("全体进攻")
	for _, f := range soldiers {
		//所有子类实现都调用共同的父类方法
		f.Attack()
	}
	//强调个性时
	fmt.Println("敌寇势大，法师进攻，骑士防守")
	for _, s := range soldiers {

		//类型断言方式1
		switch s.(type) {
		case *Rider:
			s.Defend()
		case *Master:
			s.Attack()
		default:
			fmt.Println("来了个什么鬼...")
		}
		//类型断言方式2
		//判断是当前s实例是不是骑士
		// 如果是:ok为true，riderPtr有值
		//如果不是：ok为false，riderPtr为nil
		if riderPtr, ok := s.(*Rider); ok {
			riderPtr.Defend()
		}
		//判断是当前s实例是不是法师
		// 如果是:ok为true，masterPtr有值
		//如果不是：ok为false，masterPtr为nil
		if masterPtr, ok := s.(*Master); ok {
			masterPtr.Attack()
		}

	}
}

//战士接口
type Soldier interface {
	Attack() (bloodLoss int)
	Defend()
}

//骑兵：战士的实现形态一
type Rider struct{}

//骑兵的进攻与防守
func (r *Rider) Attack() (bloodLoss int) {
	fmt.Println("发动战争践踏，撞死你，踩死你，扎死你")
	return 50
}
func (r *Rider) Defend() {
	fmt.Println("不能跑，我得快点走")
}

//法师：战士的实现形态二
type Master struct{}

//法师的进攻与防守
func (r *Master) Attack() (bloodLoss int) {
	fmt.Println("天崩地裂，电闪雷鸣，群星坠落，万籁俱寂，遍布下去了...")
	return 10000
}
func (r *Master) Defend() {
	fmt.Println("回城")
}

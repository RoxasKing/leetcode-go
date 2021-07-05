package main

// Tags:
// Hash
// DFS

type ThroneInheritance struct {
	child map[string][]string
	death map[string]bool
	king  string
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{
		child: map[string][]string{},
		death: map[string]bool{},
		king:  kingName,
	}
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	this.child[parentName] = append(this.child[parentName], childName)
}

func (this *ThroneInheritance) Death(name string) {
	this.death[name] = true
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	out := []string{}
	var dfs func(string)
	dfs = func(name string) {
		if !this.death[name] {
			out = append(out, name)
		}
		for _, childName := range this.child[name] {
			dfs(childName)
		}
	}
	dfs(this.king)
	return out
}

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */

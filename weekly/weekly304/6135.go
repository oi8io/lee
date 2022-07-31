package weekly304

func longestCycle(edges []int) int {
	n := len(edges)
	vis := make(map[int]bool)
	var ans = -1
Loop:
	for i := 0; i < n; i++ {
		if vis[i] {
			continue
		}
		cur := i
		var cycle = make([]int, 0)
		for !vis[cur] {
			vis[cur] = true
			cycle = append(cycle, cur)
			if cur == -1 {
				cycle = make([]int, 0)
				continue Loop
			}
			cur = edges[cur]
		}
		for k := 0; k < len(cycle); k++ {
			if cycle[k] == cur {
				l := len(cycle) - k
				if l > ans {
					ans = l
				}
			}
		}
	}
	return ans
}

/**
https://www.cnblogs.com/lfri/p/15758120.html
  int DirectMaxCycle(vector<int>& favorite) {
      int n = favorite.size();
      vector<bool> vis(n, false);
      int max_cycle = 0;
      for(int i = 0;i < n;i++) {
          if(vis[i]) continue;
          int cur = i;
          vector<int> cycle;
          while(!vis[cur]) {
              vis[cur] = true;
              cycle.push_back(cur);
              cur = favorite[cur];
          }
          for(int j = 0;j < cycle.size();j++) {
              if(cycle[j] == cur) {
                  int len = cycle.size() - j;
                  if(len > max_cycle) max_cycle = len;
                  break;
              }
          }
      }
      return max_cycle;
  }
*/

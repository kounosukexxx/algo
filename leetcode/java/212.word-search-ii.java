/*
 * @lc app=leetcode id=212 lang=java
 *
 * [212] Word Search II
 */
// 計算量バカ
import java.util.*;

// @lc code=start
class Solution {
	int height, width;

	public List<String> findWords(char[][] board, String[] words) {
		height = board.length;
		width = board[0].length;
		ArrayList<String> ans = new ArrayList<>();

		// 3 * 10^4
		for (String word : words) {
			// r * c
			LOOP: for (int r = 0; r < height; r++) {
				for (int c = 0; c < width; c++) {
					Set<Map.Entry<Integer, Integer>> visited = new HashSet<>();
					// 10 * 10 * 10くらい (菱形*10)
					if (dfs(board, visited, word, r, c)) {
						ans.add(word);
						break LOOP;
					}
				}
			}
		}

		return ans;
	}

	boolean dfs(char[][] board, Set<Map.Entry<Integer, Integer>> visited, String target, int r, int c) {
		// deep copy
		Set<Map.Entry<Integer, Integer>> visited2 = new HashSet<>();
		visited2.addAll(visited);

		if (target.length() == 0) {
			return true;
		}

		if (r < 0 || r >= height || c < 0 || c >= width) {
			return false;
		}

		if (visited2.contains(Map.entry(r, c))) {
			return false;
		}

		visited2.add(Map.entry(r, c));

		if (board[r][c] == target.charAt(0)) {
			return dfs(board, visited2, target.substring(1), r - 1, c) ||
					dfs(board, visited2, target.substring(1), r + 1, c) ||
					dfs(board, visited2, target.substring(1), r, c - 1) ||
					dfs(board, visited2, target.substring(1), r, c + 1);
		}

		return false;
	}
}
// @lc code=end

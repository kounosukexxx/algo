package l733;
/*
 * @lc app=leetcode id=733 lang=java
 *
 * [733] Flood Fill
 */

// @lc code=start
class Solution {
    public int[][] floodFill(int[][] image, int sr, int sc, int color) {
        boolean[][] visited = new boolean[image.length][image[0].length];
        fill(image, visited, sr, sc, color, image[sr][sc]);
        return image;
    }

    public void fill(int[][] image, boolean[][] visited, int r, int c, int toColor, int fromColor) {
        if (r < 0 || r >= image.length || c < 0 || c >= image[0].length || visited[r][c]) {
            return;
        }
        visited[r][c] = true;

        if (image[r][c] == fromColor) {
            image[r][c] = toColor;
            fill(image, visited, r - 1, c, toColor, fromColor);
            fill(image, visited, r + 1, c, toColor, fromColor);
            fill(image, visited, r, c - 1, toColor, fromColor);
            fill(image, visited, r, c + 1, toColor, fromColor);
        }
    }
}
// @lc code=end

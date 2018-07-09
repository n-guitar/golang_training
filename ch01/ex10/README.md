# 以下のサイトで実行
```bash
## 1回目
$ ./fetchall https://www.amazon.co.jp/ https://www.ponparemall.com/ https://www.rakuten.co.jp/ https://www.alibaba.com/
0.87s  593748 https://www.ponparemall.com/
0.98s   96530 https://www.alibaba.com/
1.16s  469706 https://www.amazon.co.jp/
5.51s  373533 https://www.rakuten.co.jp/
5.51s elapsed

## 2回目
$ ./fetchall https://www.amazon.co.jp/ https://www.ponparemall.com/ https://www.rakuten.co.jp/ https://www.alibaba.com/
0.29s  593748 https://www.ponparemall.com/
0.66s  483247 https://www.amazon.co.jp/
0.72s   96525 https://www.alibaba.com/
5.29s  373533 https://www.rakuten.co.jp/
5.29s elapsed

## 3回目
$ ./fetchall https://www.amazon.co.jp/ https://www.ponparemall.com/ https://www.rakuten.co.jp/ https://www.alibaba.com/
0.84s  593748 https://www.ponparemall.com/
0.96s   96530 https://www.alibaba.com/
1.11s  469903 https://www.amazon.co.jp/
5.44s  373532 https://www.rakuten.co.jp/
5.44s elapsed

## 4回目
$ ./fetchall https://www.amazon.co.jp/ https://www.ponparemall.com/ https://www.rakuten.co.jp/ https://www.alibaba.com/
0.29s  593748 https://www.ponparemall.com/
0.67s  482334 https://www.amazon.co.jp/
0.74s   96525 https://www.alibaba.com/
5.31s  373532 https://www.rakuten.co.jp/
5.32s elapsed

## 5回目
$ ./fetchall https://www.amazon.co.jp/ https://www.ponparemall.com/ https://www.rakuten.co.jp/ https://www.alibaba.com/
0.66s  461932 https://www.amazon.co.jp/
0.95s  593748 https://www.ponparemall.com/
1.11s   96530 https://www.alibaba.com/
5.27s  373532 https://www.rakuten.co.jp/
5.27s elapsed
```
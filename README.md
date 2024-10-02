# fonthas

ttfフォントが対応していない文字を、CSVの文字列から検出します。  
1行目のデータを各カラム名として扱います。

## Usage
```
fonthas <ttf path> <csv path>
```

## Example
example.csv
```
Team,Genre,Artist,Title
AESTORIA,Dramatical,Lime,Luminaria
71,D区少女 -Girls in Disco Zone-,Retro Disco,T2o Feat 盧静,Disco☆Girl(蹦迪少女)
GOMO,Futurecore,石田愛翔(F Rabbeat)(mov:𝗬𝗘𝗔𝗛𝗧𝗨𝗡𝗔),U ARE
```

コマンドライン
```
> fonthas GenShinGothic-Bold.ttf example.csv
2[Title]: 蹦 in Disco☆Girl(蹦迪少女)
3[Artist]: 𝗬 in 石田愛翔(F Rabbeat)(mov:𝗬𝗘𝗔𝗛𝗧𝗨𝗡𝗔)
3[Artist]: 𝗘 in 石田愛翔(F Rabbeat)(mov:𝗬𝗘𝗔𝗛𝗧𝗨𝗡𝗔)
3[Artist]: 𝗔 in 石田愛翔(F Rabbeat)(mov:𝗬𝗘𝗔𝗛𝗧𝗨𝗡𝗔)
3[Artist]: 𝗛 in 石田愛翔(F Rabbeat)(mov:𝗬𝗘𝗔𝗛𝗧𝗨𝗡𝗔)
3[Artist]: 𝗧 in 石田愛翔(F Rabbeat)(mov:𝗬𝗘𝗔𝗛𝗧𝗨𝗡𝗔)
3[Artist]: 𝗨 in 石田愛翔(F Rabbeat)(mov:𝗬𝗘𝗔𝗛𝗧𝗨𝗡𝗔)
3[Artist]: 𝗡 in 石田愛翔(F Rabbeat)(mov:𝗬𝗘𝗔𝗛𝗧𝗨𝗡𝗔)
3[Artist]: 𝗔 in 石田愛翔(F Rabbeat)(mov:𝗬𝗘𝗔𝗛𝗧𝗨𝗡𝗔)
```
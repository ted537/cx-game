package ui;

type CharData struct {
	w,h int
	tx,ty int
	ascii int
	index int
}

var charDatas = []CharData {
	{6,30,0,0,' ',0},
	{11,30,6,0,'!',1},
	{11,30,17,0,'"',2},
	{15,30,28,0,'#',3},
	{13,30,43,0,'$',4},
	{15,30,56,0,'%',5},
	{15,30,71,0,'&',6},
	{6,30,86,0,'\'',7},
	{9,30,92,0,'(',8},
	{9,30,101,0,')',9},
	{17,30,110,0,'*',10},
	{13,30,127,0,'+',11},
	{6,30,140,0,',',12},
	{13,30,146,0,'-',13},
	{6,30,159,0,'.',14},
	{13,30,165,0,'/',15},
	{13,30,178,0,'0',16},
	{13,30,191,0,'1',17},
	{13,30,204,0,'2',18},
	{13,30,217,0,'3',19},
	{13,30,230,0,'4',20},
	{13,30,243,0,'5',21},
	{13,30,0,30,'6',22},
	{13,30,13,30,'7',23},
	{13,30,26,30,'8',24},
	{13,30,39,30,'9',25},
	{6,30,52,30,':',26},
	{6,30,58,30,';',27},
	{11,30,64,30,'<',28},
	{13,30,75,30,'=',29},
	{11,30,88,30,'>',30},
	{13,30,99,30,'?',31},
	{13,30,112,30,'@',32},
	{13,30,125,30,'A',33},
	{13,30,138,30,'B',34},
	{13,30,151,30,'C',35},
	{13,30,164,30,'D',36},
	{13,30,177,30,'E',37},
	{13,30,190,30,'F',38},
	{13,30,203,30,'G',39},
	{13,30,216,30,'H',40},
	{13,30,229,30,'I',41},
	{13,30,242,30,'J',42},
	{13,30,0,60,'K',43},
	{13,30,13,60,'L',44},
	{15,30,26,60,'M',45},
	{13,30,41,60,'N',46},
	{13,30,54,60,'O',47},
	{13,30,67,60,'P',48},
	{13,30,80,60,'Q',49},
	{13,30,93,60,'R',50},
	{13,30,106,60,'S',51},
	{13,30,119,60,'T',52},
	{13,30,132,60,'U',53},
	{13,30,145,60,'V',54},
	{15,30,158,60,'W',55},
	{13,30,173,60,'X',56},
	{13,30,186,60,'Y',57},
	{13,30,199,60,'Z',58},
	{9,30,212,60,'[',59},
	{13,30,221,60,'\\',60},
	{9,30,234,60,']',61},
	{13,30,243,60,'^',62},
	{9,30,0,90,'_',63},
	{8,30,9,90,'`',64},
	{13,30,17,90,'a',65},
	{13,30,30,90,'b',66},
	{13,30,43,90,'c',67},
	{13,30,56,90,'d',68},
	{13,30,69,90,'e',69},
	{13,30,82,90,'f',70},
	{13,30,95,90,'g',71},
	{13,30,108,90,'h',72},
	{13,30,121,90,'i',73},
	{13,30,134,90,'j',74},
	{13,30,147,90,'k',75},
	{13,30,160,90,'l',76},
	{15,30,173,90,'m',77},
	{13,30,188,90,'n',78},
	{13,30,201,90,'o',79},
	{13,30,214,90,'p',80},
	{13,30,227,90,'q',81},
	{13,30,240,90,'r',82},
	{13,30,0,120,'s',83},
	{13,30,13,120,'t',84},
	{13,30,26,120,'u',85},
	{13,30,39,120,'v',86},
	{15,30,52,120,'w',87},
	{13,30,67,120,'x',88},
	{13,30,80,120,'y',89},
	{13,30,93,120,'z',90},
	{11,30,106,120,'{',91},
	{6,30,117,120,'|',92},
	{11,30,123,120,'}',93},
	{15,30,134,120,'~',94},
};

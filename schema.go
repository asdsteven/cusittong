package main

type careerS struct {
	slug, en, ch string
	courses      []int
}

type termS struct {
	slug, en, ch string
	classes      []int
}

type careerTermS struct {
	career   int
	term     int
	subjects []int
}

type subjectS struct {
	slug, en string
	courses  []int
}

type teacherS struct {
	name    string
	classes []int
}

type courseS struct {
	career   int
	subject  int
	code     string
	title    string
	units    string
	reserves []reserveS
	sections []int
}

type reserveS struct {
	major         string
	quota, enroll int
}

type sectionS struct {
	course  int
	code    string
	classes []int
}

type classS struct {
	term      int
	course    int
	nbr       string
	quota     int
	vacancy   int
	component string
	section   string
	language  string
	teachers  []int
	meetings  []int
	add, drop bool
	dept      string
}

type meetingS struct {
	class  int
	period *string
	room   *string
	date   *string
}

type databaseS struct {
	careers     []careerS
	terms       []termS
	careerTerms []careerTermS
	latestTerms int
	subjects    []subjectS
	teachers    []teacherS
	courses     []courseS
	classes     []classS
	meetings    []meetingS
}

var db = databaseS{
	careers: []careerS{
		{"PGDE", "Postgraduate - PGDE", "學位教師教育文憑生", nil},
		{"RPG", "Postgraduate - Research", "研究式課程研究生", nil},
		{"TPG", "Postgraduate - Taught", "修讀式課程研究生", nil},
		{"UG", "Undergraduate", "本科生", nil},
	},
}

type rowHeadS struct {
	code     string
	nbr      string
	title    string
	units    string
	teachers string
	reserves []reserveS
	rowBody  []rowBodyS
}

type rowBodyS struct {
	quota     int
	vacancy   int
	component string
	section   string
	language  string
	rowFoot   []rowFootS
	dept      string
}

type rowFootS struct {
	period *string
	room   *string
	date   *string
	add    bool
	drop   bool
}

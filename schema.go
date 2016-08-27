package main

type careerS struct {
	slug, en, ch string
	courses      []int
}

type termS struct {
	slug, en, ch string
	groups       []int
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

type courseS struct {
	career  int
	subject int
	code    string
	title   string
	units   string
	groups  []int
}

type groupS struct {
	course   int
	term     int
	code     string
	reserves []reserveS
	classes  []int
}

type reserveS struct {
	major  string
	quota  int
	enroll int
}

type teacherS struct {
	name    string
	classes []int
}

type classS struct {
	group     int
	nbr       string
	quota     int
	vacancy   int
	component string
	section   string
	language  string
	teachers  []int
	meetings  []int
	add       bool
	drop      bool
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
	courses     []courseS
	groups      []groupS
	teachers    []teacherS
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
	row      int
	code     string
	group    string
	nbr      string
	title    string
	units    string
	teachers []string
	reserves bool
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

type courseDetailS struct {
	title       string
	units       string
	grading     string
	components  [][2]string
	scheduled   bool
	addConsent  string
	dropConsent string
	requirement string
	attribute   string
	description string
}

type courseMeetingS struct {
	day     int
	start   string
	end     string
	room    string
	teacher string
	dates   string
}

type courseSectionS struct {
	section  string
	status   int
	meetings []courseMeetingS
}

type courseSectionsS struct {
	terms    [][2]string
	term     int
	sections []courseSectionS
	more     bool
}

type fragmentS struct {
	course   []courseS
	groups   []groupS
	teachers []teacherS
	classes  []classS
	meetings []meetingS
}

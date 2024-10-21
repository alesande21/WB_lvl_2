package solver

import "testing"

const succeed = "\u2713"
const failed = "\u2717"

func TestSolution_StringUnpacking(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Пустая строка",
			args:    args{},
			want:    "",
			wantErr: false,
		},
		{
			name:    "Один символ без повторения",
			args:    args{"a"},
			want:    "a",
			wantErr: false,
		},
		{
			name:    "Один символ c повторением",
			args:    args{"у4"},
			want:    "уууу",
			wantErr: false,
		},
		{
			name:    "Неправильый ввод с числа вначале строки",
			args:    args{"4у"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Неправильый ввод с числа в середине",
			args:    args{"f24у"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Множество разных символов с повторением",
			args:    args{"f2у4"},
			want:    "ffуууу",
			wantErr: false,
		},
		{
			name:    "Множество символов с повторением",
			args:    args{"a4bc2d5e"},
			want:    "aaaabccddddde",
			wantErr: false,
		},
		{
			name:    "Escape-последовательност выводит числа",
			args:    args{"qwe\\4\\5"},
			want:    "qwe45",
			wantErr: false,
		},
		{
			name:    "Escape-последовательност выводит числа с повторением",
			args:    args{"qwe\\45"},
			want:    "qwe44444",
			wantErr: false,
		},
		{
			name:    "Escape-последовательност выводит \\ с повторением",
			args:    args{"qwe\\\\5"},
			want:    "qwe\\\\\\\\\\",
			wantErr: false,
		},
	}

	/*
		qwe\4\5 => qwe45 (*)
		qwe\45 => qwe44444 (*)
		qwe\\5 => qwe\\\\\ (*)
	*/
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Solution{}
			got, err := s.StringUnpacking(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("\t%s\tТест %s: Неккоректная строка ввода %v", failed, tt.name, tt.args)
				return
			}
			if got != tt.want {
				t.Errorf("\t%s\tТест %s: StringUnpacking() = %v, want %v", failed, tt.name, got, tt.want)
			} else {
				t.Logf("\t%s\tТест %s:\tПройден!", succeed, tt.name)
			}
		})
	}
}

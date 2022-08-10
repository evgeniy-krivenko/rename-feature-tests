package flags

import "flag"

type flags struct {
	AbsPath            *string
	TagPrefix          *string
	ExcludedFolderName *string
	TargetFolderName   *string
}

func (f *flags) parseFlags() *flags {
	f.AbsPath = flag.String("path", "", "absolute path for your stream folder")
	f.TagPrefix = flag.String("tag", "", "tag for checks or add")
	f.ExcludedFolderName = flag.String("exf", "", "excluded folder name")
	f.TargetFolderName = flag.String("target", "", "target folder")
	flag.Parse()
	return f
}

func GetFlags() *flags {
	f := flags{}
	return f.parseFlags()
}

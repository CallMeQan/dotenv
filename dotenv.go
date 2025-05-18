package dotenv

// +====================+
// |   Public function  |
// +====================+

// Load single file or multiple files.
//
// Better call this function as close as start of program.
//
//
func Load(filenames ...string) error {
	filenames = getFilenamesOrDefault(filenames)

	for _, filename := range filenames {
		if err := loadFile(filename, false); err != nil {
			return err
		}
	}
	return nil
}

// Overwrite any env variables with input file(s).
//
// Better call this function as close as start of program.
//
// Example:
//
//		dotenv.ForceLoadFromFiles(".proc.env", ".dev.env")
// 		// Priorty: .dev.env > .proc.env > default variables
func ForceLoad(filenames ...string) error {
	filenames = getFilenamesOrDefault(filenames)

	for _, filename := range filenames {
		if err := loadFile(filename, true); err != nil {
			return err
		}
	}
	return nil
}

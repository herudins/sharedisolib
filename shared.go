package shared

// Info for internal use
// Changelog
// 3.0 - next development from source github.com/iyusa/shared
// 2.4 - bug fix tool.StrToFloat, 32 => 64
// 2.3 - add validate mti on load
// 2.2 - add conn.Close() to handler
// 2.1 - add bit 9
// 2.0 - tool.WordWraps
// 1.9 - iso Print
// 1.8 - Rename TrannsactionHandler -> ExecuteHandler
// 1.7 - add server v2, add PcReversal, CustomError
// 1.6 - add md5
// 1.5 - add fixed string for pln
// 1.3 - add bit 60 to iso
func Info() string {
	return "Version 3.0 - Herudin Saifuloh - herudinsaifuloh@gmail.com"
}



var Configuration = map[string]PathConfig{
	"bash":    Allowed,
	"bison":   Allowed,  // added
	"brotli":  Allowed,	 // added
	"ccache":  Allowed,
	"dd":      Allowed,
	"diff":    Allowed,
	"dlv":     Allowed,
	"expr":    Allowed,
	"flex":    Allowed,   // added
	"fuser":   Allowed,
	"getopt":  Allowed,
	"git":     Allowed,
	"hexdump": Allowed,
	"jar":     Allowed,
	"java":    Allowed,
	"javap":   Allowed,
	"lsof":    Allowed,
	"m4":      Allowed,   // added
	"nproc":   Allowed,	  // added
	"openssl": Allowed,
	"pstree":  Allowed,
	"rsync":   Allowed,
	"sh":      Allowed,
	"tr":      Allowed,
	"unzip":   Allowed,
	"zip":     Allowed,
	"arm-linux-androidkernel-as":     Allowed,
	"arm-linux-androidkernel-ld":     Allowed,

	// Host toolchain is removed. In-tree toolchain should be used instead.
	// GCC also can't find cc1 with this implementation.
	"ar":         Forbidden,
	"as":         Forbidden,
	"cc":         Forbidden,
	"clang":      Forbidden,
	"clang++":    Forbidden,
	"gcc":        Forbidden,
	"g++":        Forbidden,
	"ld":         Forbidden,
	"ld.bfd":     Forbidden,
	"ld.gold":    Forbidden,
	"pkg-config": Forbidden,

	// These are toybox tools that only work on Linux.
	"pgrep": LinuxOnlyPrebuilt,
	"pkill": LinuxOnlyPrebuilt,
	"ps":    LinuxOnlyPrebuilt,
}


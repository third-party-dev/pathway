package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

func is_decimal(arg string) bool {
	for _, c := range(arg) {
		if (c < '0' || c > '9') {
			return false
		}
	}
	return true;
}

func is_hex(arg string) bool {
	// TODO: Is X valid?
	if (len(arg) < 3) || (arg[0] != '0') || (arg[1] != 'x' && arg[1] != 'X') {
		return false
	}

	i := 2
	for ; i < len(arg); i++ {
		if (arg[i] < '0' || arg[i] > '9') && (arg[i] < 'a' && arg[i] > 'f') && (arg[i] < 'A' && arg[i] > 'F') {
			return false
		}
	}
	return true;
}

func is_octal(arg string) bool {
	// TODO: Is X valid?
	if (len(arg) < 2) || (arg[0] != '0') {
		return false
	}

	i := 1
	for ; i < len(arg); i++ {
		if (arg[i] < '0' || arg[i] > '7') {
			return false
		}
	}
	return true;
}

func is_valid_index(arg string) bool {
	return (is_decimal(arg) || is_hex(arg) || is_octal(arg))
}


func way_count_elems(path string) uint64 {
	var count uint64 = 0

	if len(path) == 0 {
		return 0
	}

	for i := 0; i < len(path); i++ {
		if (path[i] == ':') {
			count += 1
		}
	}

	// return count > 0 ? count + 1 : 1
	if (count > 0) {
		return count + 1
	}
	return 1
}

func _insert_elem(mode int, dst *string, path string, idx uint64, npath string) {
	var count uint64 = 0

	if (idx == 0) {
		// TODO: Implement mode 1
		if (mode == 2) {
			fmt.Printf("%s", npath)
			if (len(path) > 0) {
				fmt.Printf("%c", ':')
			}
		}
	}

	for i, w := 0, 0; i < len(path); i += w {
		/* Extract multibyte UTF-8 code points. */
		c, width := utf8.DecodeRuneInString(path[i:])
		w = width

		// TODO: Implement mode 1
		if (mode == 2) {			
			fmt.Printf("%c", c)
		}
		if (path[i] == ':') {
			count++;
			if (count == idx) {
				// TODO: Implement mode 1
				if (mode == 2) {
					fmt.Printf("%s", npath)
					if (len(path) - i > 0) {
						fmt.Printf("%c", ':')
					}
				}
			}
		}
	}

	if (count + 1 == idx) {
		// TODO: Implement mode 1
		if (mode == 2) {
			if (len(path) > 0) {
				fmt.Printf("%c", ':')
			}
			fmt.Printf("%s", npath)
		}
	}
}

func way_insert_print(path string, idx uint64, npath string) {
	_insert_elem(2, nil, path, idx, npath)
}

func _delete_elem(mode int, dst *string, path string, idx uint64) {

    var count uint64 = 0
    
    //dst_idx := 0

	for i, w := 0, 0; i < len(path); i += w {
		/* Extract multibyte UTF-8 code points. */
		c, width := utf8.DecodeRuneInString(path[i:])
		w = width
    //for i := 0; i < len(path); i++ {
        if (path[i] == ':') {
            count++
            if (idx == 0 && count == 1) {
				continue
			}
            if (idx == count) {
				continue
			}
        }
        if (idx != count) {
            // if (mode == 1) {
            //     if (dst) dst[dst_idx++] = path[i];
            //     if (dst_len) *dst_len++;
            // }
            if (mode == 2) {
				fmt.Printf("%c", c)
			}
        }
    }
}

func way_delete_print(path string, idx uint64) {
	_delete_elem(2, nil, path, idx)
}

func _get_elem(mode int, dst *string, path string, idx uint64) {
    var count uint64 = 0
    
    //size_t dst_idx = 0;

	for i, w := 0, 0; i < len(path); i += w {
		/* Extract multibyte UTF-8 code points. */
		c, width := utf8.DecodeRuneInString(path[i:])
		w = width
    
        if (path[i] == ':') {
			count++
		}
        if (path[i] == ':' && idx == count) {
			continue
		}
        if (idx == count) {
            // if (mode == 1) {
            //     if (dst) dst[dst_idx++] = path[i];
            //     if (dst_len) *dst_len++;
            // }
            if (mode == 2) {
                fmt.Printf("%c", c);
            }
        }
    }
}

func way_get_print(path string, idx uint64) {
	_get_elem(2, nil, path, idx)
}


func process_insert(command string, path string, args []string) {

	var idx uint64 = 0

	index_given := 0
	
	i := 0
	for ; i < len(args); i++ {
		if args[i] == "--" {
			i++
			break
		} else if args[i][0] == '-' {
			if args[i] == "--help" {
				goto usage
			} else if (args[i] == "-i" || args[i] == "--index") && index_given == 0 {
				index_given = 1
				if i+1 < len(args) {
					i++
					if is_valid_index(args[i]) {
						var errno error
						idx, errno = strconv.ParseUint(args[i], 0, 64)
						// TODO: Check for overflow and underflow. (How to use limits.h in Go?)
						if /*idx == limits.LONG_MIN || idx == limits.LONG_MAX ||*/ errno != nil {
							goto usage
						}
					} else {
						goto usage
					}
				}
			} else if (args[i] == "-t" || args[i] == "--tail") && index_given == 0 {
				index_given = 1
				idx = way_count_elems(path);
			} else {
				goto usage
			}
			continue
		}
		break
	}

	if (len(args) - i == 1) {
		npath := args[i];
		way_insert_print(path, idx, npath)
		os.Exit(0)
	}

usage:
	fmt.Printf("Insert Usage: %s", path)
	os.Exit(1)
}

func process_delete(command string, path string, args []string) {
	var idx uint64 = 0

	index_given := 0

	i := 0
	for ; i < len(args); i++ {
		if args[i] == "--" {
			i++
			break
		} else if args[i][0] == '-' {
			if args[i] == "--help" {
				goto usage
			} else if (args[i] == "-i" || args[i] == "--index") && index_given == 0 {
				index_given = 1
				if i+1 < len(args) {
					i++
					if is_valid_index(args[i]) {
						var errno error
						idx, errno = strconv.ParseUint(args[i], 0, 64)
						// TODO: Check for overflow and underflow. (How to use limits.h in Go?)
						if /*idx == limits.LONG_MIN || idx == limits.LONG_MAX ||*/ errno != nil {
							goto usage
						}
					} else {
						goto usage
					}
				}
			} else if (args[i] == "-t" || args[i] == "--tail") && index_given == 0 {
				index_given = 1
				idx = way_count_elems(path);
				if idx > 0 {
					idx--
				}
			} else {
				goto usage
			}
			continue
		}
		break
	}

    if (len(args) - i == 0) {
		way_delete_print(path, idx)
		os.Exit(0)
	}

usage:
    /* Dump path even during errors to prevent an error from busting PATH */
	fmt.Printf("%s", path)
	fmt.Printf("Usage: ...")
    //delete_usage(command);
    os.Exit(1);
}

func process_get(command string, path string, args []string) {
	var idx uint64 = 0

	index_given := 0

	i := 0
	for ; i < len(args); i++ {
		if args[i] == "--" {
			i++
			break
		} else if args[i][0] == '-' {
			if args[i] == "--help" {
				goto usage
			} else if (args[i] == "-i" || args[i] == "--index") && index_given == 0 {
				index_given = 1
				if i+1 < len(args) {
					i++
					if is_valid_index(args[i]) {
						var errno error
						idx, errno = strconv.ParseUint(args[i], 0, 64)
						// TODO: Check for overflow and underflow. (How to use limits.h in Go?)
						if /*idx == limits.LONG_MIN || idx == limits.LONG_MAX ||*/ errno != nil {
							goto usage
						}
					} else {
						goto usage
					}
				}
			} else if (args[i] == "-t" || args[i] == "--tail") && index_given == 0 {
				index_given = 1
				idx = way_count_elems(path);
				if idx > 0 {
					idx--
				}
			} else {
				goto usage
			}
			continue
		}
		break
	}

    if (len(args) - i == 0) {
		way_get_print(path, idx)
		os.Exit(0)
	}

usage:
	fmt.Printf("Usage: ...")
    //get_usage(command);
    os.Exit(1);
}

func process_count(command string, path string, args []string) {

    i := 0;
    for ; i < len(args); i++ {
        if args[i][0] == '-' {
            if args[i] == "--help" {
                goto usage;
            } else {
                goto usage;
            }
            continue;
        }
        break;
    }

    if (len(args) != i) {
        goto usage;
    }

    fmt.Printf("%d", way_count_elems(path))
    os.Exit(0)

usage:
    //count_usage(command);
	fmt.Printf("Usage: ...")
    os.Exit(1);
}

func process_bytes(command string, path string, args []string) {

    i := 0;
    for ; i < len(args); i++ {
        if args[i][0] == '-' {
            if args[i] == "--help" {
                goto usage;
            } else {
                goto usage;
            }
            continue;
        }
        break;
    }

    if (len(args) != i) {
        goto usage;
    }

    fmt.Printf("%d", len(path))
    os.Exit(0)

usage:
    //count_usage(command);
	fmt.Printf("Usage: ...")
    os.Exit(1);
}

func process_chars(command string, path string, args []string) {

    count, i := 0, 0

    for ; i < len(args); i++ {
        if args[i][0] == '-' {
            if args[i] == "--help" {
                goto usage;
            } else {
                goto usage;
            }
            continue;
        }
        break;
    }

    if (len(args) != i) {
        goto usage;
    }

	// Note: range implicitly decodes runes
	for i = range path {
		count++
	}

    fmt.Printf("%d", count)
    os.Exit(0)

usage:
    //count_usage(command);
	fmt.Printf("Usage: ...")
    os.Exit(1);
}

func main() {
	var subcommand string

	path := os.Getenv("PATH")

	/*
	   Usage: <command> [options] <sub command> [options] <args>
	   States:
	    - Command
	    - Command Options
	    - Subcommand
	    - Subcommand Options
	    - Subcommand Arguments
	*/

	/* 1. Get command */
	//command := os.Args[0]

	// https://github.com/famz/SetLocale/blob/master/SetLocale.go
	// C.setlocale(0, "")
	/* 2. Process options until sub_command */
	if len(os.Args) <= 1 {
		fmt.Printf("%s", path)
		os.Exit(0)
	}

	//offset := 0
	i := 1
	for ; i < len(os.Args); i++ {
	//for i, v := range os.Args[1:] {
		if os.Args[i][0] == '-' {
			if os.Args[i] == "--help" {
				goto usage
			} else {
				goto usage
			}
		} else if subcommand == "" && os.Args[i][0] != '-' {
			subcommand = os.Args[i]
			i++
			break
		}
	}

	if subcommand != "" {
		if subcommand == "insert" {
			process_insert(os.Args[0], path, os.Args[i:])
		} else if subcommand == "delete" {
			process_delete(os.Args[0], path, os.Args[i:])
		} else if subcommand == "count" {
			process_count(os.Args[0], path, os.Args[i:])
		} else if subcommand == "bytes" {
			process_bytes(os.Args[0], path, os.Args[i:])
		} else if subcommand == "chars" {
			process_chars(os.Args[0], path, os.Args[i:])
		} else if subcommand == "get" {
			process_get(os.Args[0], path, os.Args[i:])
		}
	}

	// TODO: Consider path output options?

usage:
	fmt.Printf("%s", path)
	fmt.Printf("Usage: ...");
	//print_usage(command);
	os.Exit(1)

}

package way

import (
	"fmt"
	"unicode/utf8"
)

func Count(path string) uint64 {
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

func InsertPrint(path string, idx uint64, npath string) {
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

func DeletePrint(path string, idx uint64) {
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

func GetPrint(path string, idx uint64) {
	_get_elem(2, nil, path, idx)
}
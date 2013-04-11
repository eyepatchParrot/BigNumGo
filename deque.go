package main

//import "log"

type Deque struct {
	buffer    []uint32
	sz_buffer int
	sz_deque  int
	ix_start  int
}

// GetAt(ix_deque) val
// SetAt(ix_deque, val)
// PushBack(val)
// Size() int
// String() string
// dequeIdxToBufferIdx(ix_deque) ix_buffer
// reserveSpaceFor(numMore) err

func (d *Deque) GetAt(ix_deque int) (val uint32) {
	if d.buffer == nil {
		panic(InvalidSliceError("buffer", "GetAt"))
	}

	if ix_deque >= d.sz_deque {
		panic(OutOfBoundsError("GetAt"))
	}
	ix_buffer := d.dequeIdxToBufferIdx(ix_deque)
	return d.buffer[ix_buffer]
}

func (d *Deque) SetAt(ix_deque int, val uint32) {
	if d.buffer == nil {
		panic(InvalidSliceError("d.buffer", "SetAt"))
	}

	if ix_deque >= d.sz_deque {
		panic(OutOfBoundsError("GetAt"))
	}
	ix_buffer := d.dequeIdxToBufferIdx(ix_deque)
	d.buffer[ix_buffer] = val
}

func (d *Deque) PushBack(val uint32) {
	d.reserveSpaceFor(1)
	d.sz_deque++
	d.SetAt(d.sz_deque-1, val)
}

func (d *Deque) Size() int {
	return d.sz_deque
}

func (d *Deque) String() (ret string) {
	for i := 0; i < d.Size(); i++ {
		str := UIntToHexString(d.GetAt(i), false)
		if i < d.Size() - 1 {
			ret += str + " "
		} else {
			ret += str
		}
	}
	return ret
}

func (d *Deque) dequeIdxToBufferIdx(ix_deque int) (ix_buffer int) {
	ix_buffer = ix_deque + d.ix_start
	if ix_buffer >= d.sz_buffer {
		ix_buffer -= d.sz_buffer
	}
	return ix_buffer
}

func (d *Deque) reserveSpaceFor(numMore int) {
	if (d.sz_deque + numMore) <= d.sz_buffer {
		return
	}
	newSize := d.sz_buffer
	if d.sz_buffer == 0 {
		newSize = 16
	}
	for newSize < d.sz_deque + numMore {
		newSize *= 2
	}
	newBuffer := make([]uint32, newSize)
//	log.Println(newSize, newBuffer)
	for i := 0; i < d.sz_deque; i++ {
		newBuffer[i] = d.GetAt(i)
	}
	d.buffer = newBuffer
	d.ix_start = 0
	d.sz_buffer = newSize
}

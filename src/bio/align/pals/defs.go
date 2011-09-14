package pals
// Copyright ©2011 Dan Kortschak <dan.kortschak@adelaide.edu.au>
//
//   This program is free software: you can redistribute it and/or modify
//   it under the terms of the GNU General Public License as published by
//   the Free Software Foundation, either version 3 of the License, or
//   (at your option) any later version.
//
//   This program is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//   GNU General Public License for more details.
//
//   You should have received a copy of the GNU General Public License
//   along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
import (
	"bio/feat"
	"bio/seq"
)

const (
	MaxIGap    = 5
	DiffCost   = 3
	SameCost   = 1
	MatchCost  = DiffCost + SameCost
	BlockCost  = DiffCost * MaxIGap
	RMatchCost = DiffCost + 1
)

var (
	binSize    int = 1 << 10
	minPadding int = 50
)

type FeaturePair struct {
	A, B   *feat.Feature
	Score  int
	Error  float64
	Strand int8
}

type Contig struct {
	seq  *seq.Seq
	from int
}

type SeqMap struct {
	contigs []Contig
	binMap  []int
}

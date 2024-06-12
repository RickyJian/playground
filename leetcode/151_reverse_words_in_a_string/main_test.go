package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	var tests = []*struct {
		s        string
		expected string
	}{
		//{
		//	s:        "the sky is blue",
		//	expected: "blue is sky the",
		//},
		{
			s:        "  hello world  ",
			expected: "world hello",
		},
		{
			s:        "a good   example",
			expected: "example good a",
		},
		{
			s:        "EPY2giL",
			expected: "EPY2giL",
		},
		{
			s:        " s       pokEfM9Smb     CkB   I YR6mlSb  NMUjwXO1zT      1jYFIj5     3yKuOFv   VUgEc     SDj4      X9 yxJ7m1MFQ3     2czWYkN   mcn7t UQB      0uz   gWR  6cqWMJTboG  kc     qvpXw   Vemp     ckiI8Gnj GDgyIK7e    kJ       HNuW      G6axXAKZ     mrdlIlUz    1ei  4uriBAE       uBq     iP3V       tgn1 tcp4    OQkssUH3     Cu       JRvD YUUr0K   5NfLPJLii    EBWEJ    99q    Gw    f AkEwlaJJ12  NjygVYQxZt       o GWxAguB    gBxLKLnO   Fz2y    B   Exih2X  Ob9Z6m8o  Ko   t6      QBR     a1Lq  AU35   QN  UIRFK   PEOn4      5JQvOh jmOL42    Xd9lK1  Rmh D    PzSeb   MPh     7 b5tC6     niT    54w3VPhy  h3w5esv  5       B     gUQ     Ggi2T 4IIX84       MUfe      sgzWdNZ      yYIKbONI8U    uvZOd4d  ne0oMp     LJTF       p  6       Tu6BaSwM AUFHHl       H    kimskH9s      VdwSz     b  6M xVeFv6DNN       5M64     aT9x5LF9     vO   epjgc     isAROwj2A      0hAFVXt  nyZd5ISv  XKcRG      71J    7e9K 4P9dpRW       VAfA Rmy4M9sF       GX7Cv  k8      8yilzwDHO   J   Ljq3C      CP5096AS       Fw8  0FARKi6x1v      LUT5UeJeU  GBjxE LpF1cyNa   X9ceJx  Q2YX3r dMpdX       K Ij  Q2cOu     wg   PBvcVYz  93zgWnXC       0u   of7mdXO     I6xRvT     Ao Y9e   4l  TVi   s    1tckz5O     oJAM8y bI6ppw    h   ICtUeokj  hNooLvzq Edx1I      o   yJ0ebu   V7 60h      clMtNMYjx yuIxtWb     6    94YHa8c    YT1aXn  je06     6xV5eVX7ug  z6r   U1   6x  ss  0PZAfLVjEi  pdZJmxA9   Qy17QwX MA     Nv   j Fap3X    Lme1mm rlusXgLD     f7aX0mEs0n   RkB    tA       h8   Dli FxsVvNgq  qAb8qUO",
			expected: "qAb8qUO FxsVvNgq Dli h8 tA RkB f7aX0mEs0n rlusXgLD Lme1mm Fap3X j Nv MA Qy17QwX pdZJmxA9 0PZAfLVjEi ss 6x U1 z6r 6xV5eVX7ug je06 YT1aXn 94YHa8c 6 yuIxtWb clMtNMYjx 60h V7 yJ0ebu o Edx1I hNooLvzq ICtUeokj h bI6ppw oJAM8y 1tckz5O s TVi 4l Y9e Ao I6xRvT of7mdXO 0u 93zgWnXC PBvcVYz wg Q2cOu Ij K dMpdX Q2YX3r X9ceJx LpF1cyNa GBjxE LUT5UeJeU 0FARKi6x1v Fw8 CP5096AS Ljq3C J 8yilzwDHO k8 GX7Cv Rmy4M9sF VAfA 4P9dpRW 7e9K 71J XKcRG nyZd5ISv 0hAFVXt isAROwj2A epjgc vO aT9x5LF9 5M64 xVeFv6DNN 6M b VdwSz kimskH9s H AUFHHl Tu6BaSwM 6 p LJTF ne0oMp uvZOd4d yYIKbONI8U sgzWdNZ MUfe 4IIX84 Ggi2T gUQ B 5 h3w5esv 54w3VPhy niT b5tC6 7 MPh PzSeb D Rmh Xd9lK1 jmOL42 5JQvOh PEOn4 UIRFK QN AU35 a1Lq QBR t6 Ko Ob9Z6m8o Exih2X B Fz2y gBxLKLnO GWxAguB o NjygVYQxZt AkEwlaJJ12 f Gw 99q EBWEJ 5NfLPJLii YUUr0K JRvD Cu OQkssUH3 tcp4 tgn1 iP3V uBq 4uriBAE 1ei mrdlIlUz G6axXAKZ HNuW kJ GDgyIK7e ckiI8Gnj Vemp qvpXw kc 6cqWMJTboG gWR 0uz UQB mcn7t 2czWYkN yxJ7m1MFQ3 X9 SDj4 VUgEc 3yKuOFv 1jYFIj5 NMUjwXO1zT YR6mlSb I CkB pokEfM9Smb s",
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, reverseWords(test.s))
	}
}

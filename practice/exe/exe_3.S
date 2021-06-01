TEXT main.main(SB) /Users/bakatora/ws/gom/playg/practice/exe/exe.go
  exe.go:6		0x1061020		65488b0c2530000000	MOVQ GS:0x30, CX			
  exe.go:6		0x1061029		488d4424e8		LEAQ -0x18(SP), AX			
  exe.go:6		0x106102e		483b4110		CMPQ 0x10(CX), AX			
  exe.go:6		0x1061032		0f86e8000000		JBE 0x1061120				
  exe.go:6		0x1061038		4881ec98000000		SUBQ $0x98, SP				
  exe.go:6		0x106103f		4889ac2490000000	MOVQ BP, 0x90(SP)			
  exe.go:6		0x1061047		488dac2490000000	LEAQ 0x90(SP), BP			
  exe.go:7		0x106104f		48c744244801000000	MOVQ $0x1, 0x48(SP)			
  exe.go:8		0x1061058		48c744244002000000	MOVQ $0x2, 0x40(SP)			
  exe.go:10		0x1061061		488b442448		MOVQ 0x48(SP), AX			
  exe.go:10		0x1061066		48890424		MOVQ AX, 0(SP)				
  exe.go:10		0x106106a		48c744240802000000	MOVQ $0x2, 0x8(SP)			
  exe.go:10		0x1061073		e8c8000000		CALL main.max(SB)			
  exe.go:10		0x1061078		488b442410		MOVQ 0x10(SP), AX			
  exe.go:10		0x106107d		4889442438		MOVQ AX, 0x38(SP)			
  exe.go:11		0x1061082		48890424		MOVQ AX, 0(SP)				
  exe.go:11		0x1061086		e895f8ffff		CALL strconv.Itoa(SB)			
  exe.go:11		0x106108b		488b442408		MOVQ 0x8(SP), AX			
  exe.go:11		0x1061090		488b4c2410		MOVQ 0x10(SP), CX			
  exe.go:11		0x1061095		4889842480000000	MOVQ AX, 0x80(SP)			
  exe.go:11		0x106109d		48898c2488000000	MOVQ CX, 0x88(SP)			
  exe.go:11		0x10610a5		488d542450		LEAQ 0x50(SP), DX			
  exe.go:11		0x10610aa		48891424		MOVQ DX, 0(SP)				
  exe.go:11		0x10610ae		488d15d5660100		LEAQ go.string.*+490(SB), DX		
  exe.go:11		0x10610b5		4889542408		MOVQ DX, 0x8(SP)			
  exe.go:11		0x10610ba		48c744241005000000	MOVQ $0x5, 0x10(SP)			
  exe.go:11		0x10610c3		4889442418		MOVQ AX, 0x18(SP)			
  exe.go:11		0x10610c8		48894c2420		MOVQ CX, 0x20(SP)			
  exe.go:11		0x10610cd		e86e5dfeff		CALL runtime.concatstring2(SB)		
  exe.go:11		0x10610d2		488b442428		MOVQ 0x28(SP), AX			
  exe.go:11		0x10610d7		488b4c2430		MOVQ 0x30(SP), CX			
  exe.go:11		0x10610dc		4889442470		MOVQ AX, 0x70(SP)			
  exe.go:11		0x10610e1		48894c2478		MOVQ CX, 0x78(SP)			
  exe.go:11		0x10610e6		e835dbfcff		CALL runtime.printlock(SB)		
  exe.go:11		0x10610eb		488b442470		MOVQ 0x70(SP), AX			
  exe.go:11		0x10610f0		488b4c2478		MOVQ 0x78(SP), CX			
  exe.go:11		0x10610f5		48890424		MOVQ AX, 0(SP)				
  exe.go:11		0x10610f9		48894c2408		MOVQ CX, 0x8(SP)			
  exe.go:11		0x10610fe		6690			NOPW					
  exe.go:11		0x1061100		e83be5fcff		CALL runtime.printstring(SB)		
  exe.go:11		0x1061105		e8d6ddfcff		CALL runtime.printnl(SB)		
  exe.go:11		0x106110a		e891dbfcff		CALL runtime.printunlock(SB)		
  exe.go:12		0x106110f		488bac2490000000	MOVQ 0x90(SP), BP			
  exe.go:12		0x1061117		4881c498000000		ADDQ $0x98, SP				
  exe.go:12		0x106111e		c3			RET					
  exe.go:6		0x106111f		90			NOPL					
  exe.go:6		0x1061120		e8bb85ffff		CALL runtime.morestack_noctxt(SB)	
  exe.go:6		0x1061125		e9f6feffff		JMP main.main(SB)			
  :-1			0x106112a		cc			INT $0x3				
  :-1			0x106112b		cc			INT $0x3				
  :-1			0x106112c		cc			INT $0x3				
  :-1			0x106112d		cc			INT $0x3				
  :-1			0x106112e		cc			INT $0x3				
  :-1			0x106112f		cc			INT $0x3				
  :-1			0x1061130		cc			INT $0x3				
  :-1			0x1061131		cc			INT $0x3				
  :-1			0x1061132		cc			INT $0x3				
  :-1			0x1061133		cc			INT $0x3				
  :-1			0x1061134		cc			INT $0x3				
  :-1			0x1061135		cc			INT $0x3				
  :-1			0x1061136		cc			INT $0x3				
  :-1			0x1061137		cc			INT $0x3				
  :-1			0x1061138		cc			INT $0x3				
  :-1			0x1061139		cc			INT $0x3				
  :-1			0x106113a		cc			INT $0x3				
  :-1			0x106113b		cc			INT $0x3				
  :-1			0x106113c		cc			INT $0x3				
  :-1			0x106113d		cc			INT $0x3				
  :-1			0x106113e		cc			INT $0x3				
  :-1			0x106113f		cc			INT $0x3				


lang/textproto.ebnfù
	iso-14977
 
	



=
,
;
|
[]
{}
()
""
''
(**)
Message"Field*
Field"ScalarField"MessageFieldS
ScalarField"
field_name:"ScalarFieldValue"	FieldTerm[
MessageField"
field_name	:"MessageFieldValue"	FieldTerm3
ScalarFieldValue"
ScalarList"ScalarValue6
MessageFieldValue"MessageList"MessageValue
	FieldTerm;,U

ScalarList[53"ScalarValue,"ScalarValue]X
MessageList[75"MessageValue,"MessageValue]2
MessageValue"BraceMessage"AngleMessage+
BraceMessage{	"Message}+
AngleMessage<	"Message>Í
ScalarValue"	StringVal
"FloatVal"SignedIdentifier"
Identifier"HexSignedInteger"OctSignedInteger"DecSignedInteger"HexUnsignedInteger"OctUnsignedInteger"DecUnsignedInteger9
	StringVal"string_literal"string_literal*
FloatVal	-"float_literal$
SignedIdentifier-"ident

Identifier"ident&
HexSignedInteger-	"hex_lit&
OctSignedInteger-	"oct_lit&
DecSignedInteger-	"dec_lit
HexUnsignedInteger	"hex_lit
OctUnsignedInteger	"oct_lit
DecUnsignedInteger	"dec_lit;

field_name
"any_name"extension_name"ident/
extension_name["	type_name]D
any_name["
url_prefix/"	type_name]0
	type_name"ident."ident.

url_prefix
"url_char
"url_char6
ident"letter"letter"	dec_digit6
string_literal"double_string"single_string<
double_string"dquote	"dq_char"dquote<
single_string"squote	"sq_char"squote
dquote"
squote'0
dq_char"dq_unescaped"escape_sequence0
sq_char"sq_unescaped"escape_sequence[
dq_unescaped"dq_tab"dq_space_bang"dq_hash_bracket"dq_bracket_high
dq_tab	
dq_space_bang*
 !
dq_hash_bracket*
#[
dq_bracket_high*	
]Ùèøø[
sq_unescaped"sq_tab"sq_space_amp"sq_paren_bracket"sq_bracket_high
sq_tab	
sq_space_amp*
 &
sq_paren_bracket*
([
sq_bracket_high*	
]Ùèøø*
escape_sequence\"escaped_charª
escaped_char"	escaped_a"	escaped_b"	escaped_f"	escaped_n"	escaped_r"	escaped_t"	escaped_v"escaped_question"escaped_backslash"escaped_squote"escaped_dquote"octal_escape"
hex_escape"unicode_escape_short"unicode_escape_long
	escaped_a*
aa
	escaped_b*
bb
	escaped_f*
ff
	escaped_n*
nn
	escaped_r*
rr
	escaped_t*
tt
	escaped_v*
vv
escaped_question?
escaped_backslash\
escaped_squote'
escaped_dquote"I
octal_escape"	oct_digit(&"	oct_digit"	oct_digit>

hex_escape*
xx"	hex_digit"	hex_digitd
unicode_escape_short*
uu"	hex_digit"	hex_digit"	hex_digit"	hex_digitß
unicode_escape_long*
UU"	hex_digit"	hex_digit"	hex_digit"	hex_digit"	hex_digit"	hex_digit"	hex_digit"	hex_digit;
float_literal"float_with_suffix"int_with_suffix;
float_with_suffix"
float_body"float_suffix0
int_with_suffix	"dec_lit"float_suffixC

float_body"	dot_float"dec_dot_float"dec_exp_floatN
	dot_float."	dec_digit"	dec_digit
"exponentP
dec_dot_float	"dec_lit."	dec_digit
"exponent*
dec_exp_float	"dec_lit
"exponenth
exponent*
EE*
ee+-"	dec_digit"	dec_digit&
float_suffix*
FF*
ff;
dec_lit0'%
"digit1_9"	dec_digit6
oct_lit0"	oct_digit"	oct_digitX
hex_lit0*
XX*
xx"	hex_digit"	hex_digit.
letter*
AZ*
az*
__
	dec_digit*
09
digit1_9*
19
	oct_digit*
071
	hex_digit*
09*
AF*
afH
url_char"url_unreserved"url_sub_delim"url_pct_encodedF
url_unreserved"letter"	dec_digit-.~e
url_sub_delim!$&()*+,;=8
url_pct_encoded%"	hex_digit"	hex_digit
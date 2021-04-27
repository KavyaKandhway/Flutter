%{
#include<stdio.h>
int yyerror(char *errormsg);
%}

%%

while|if|else	 {printf("%s is a Keyword\n", yytext);}
([a-zA-Z]+[a-zA-Z0-9]*)	 {printf("%s is an Identifier\n", yytext);}
. 	{yyerror("Unknown character"); }

%%

int yywrap()
{return 1;}

int main(void)
{yylex(); return 0;}

int yyerror(char *errormsg)
{
fprintf(stderr, "%s\n", errormsg);
exit(1);
}
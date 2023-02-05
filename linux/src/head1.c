#include <stdio.h>
#include <stdlib.h>

static void do_head(FILE *f, int nlines);
int main(int argc, char *argv[])
{
	int i;
	if (argc < 2)
	{
		fprintf(stderr, "Usage: %s n [file file...]\n", argv[0]);
		exit(1);
	}
	long nlines = atol(argv[1]);
	if (argc == 2)
	{
		do_head(stdin, nlines);
	}
	else
	{
		for (i = 2; i < argc; i++)
		{

			FILE *f;
			f = fopen(argv[i], "r");
			if (!f)
			{
				perror(argv[i]);
				exit(1);
			}
			do_head(f, nlines);
			fclose(f);
		}
	}
	exit(0);
}

static void do_head(FILE *f, int nlines)
{
	int c;
	while ((c = fgetc(f)) != EOF)
	{
		if (putchar(c) < 0)
		{
			exit(1);
		}

		if (c == '\n')
		{
			nlines--;
			if (nlines == 0)
			{
				return;
			}
		}
	}
}
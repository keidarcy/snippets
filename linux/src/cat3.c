#include <stdio.h>
#include <stdlib.h>

static void do_cat(FILE *f);

int main(int argc, char *argv[])
{
	int i;

	if (argc == 1)
	{
		do_cat(stdin);
	}
	else
	{
		for (i = 1; i < argc; i++)
		{
			FILE *f;
			f = fopen(argv[i], "r");

			if (!f)
			{
				perror(argv[i]);
				exit(1);
			}

			do_cat(f);
			fclose(f);
		}
	}
	exit(0);
}

void do_cat(FILE *f)
{
	int c;

	while ((c = fgetc(f)) != EOF)
	{
		if (putchar(c) < 0)
			exit(1);
	}
}
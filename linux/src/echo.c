#include <stdio.h>

int main(int argc, char *argv[])
{
	int i;

	for (i = 1; i < argc; i++)
	{
		printf("%s%s", i == 1 ? "" : " ", argv[i]);
	}
	putchar('\n');
	return 0;
}
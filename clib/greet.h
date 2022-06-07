#ifndef __greet_h_
#define __greet_h_

#if BUILDING_LIBGREET && HAVE_VISIBILITY
#define LIBGREET_DLL_EXPORTED __attribute__((__visibility__("default")))
#else
#define LIBGREET_DLL_EXPORTED
#endif

LIBGREET_DLL_EXPORTED extern char* greet(char* name);

#endif

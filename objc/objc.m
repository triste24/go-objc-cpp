
#include "objc.h"
#import <Cocoa/Cocoa.h>
#import <Foundation/Foundation.h>

extern void Logo(char *s);

void SayObjcHello() {
    @autoreleasepool{
        NSLog(@"Say Hello From Objc");
        Logo("Say Hello From Objc From Go");
    }
}
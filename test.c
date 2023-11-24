#include <stdio.h>
#include <stdlib.h>
#include <dlfcn.h>
#include <errno.h>
#include <stdint.h>
#include <string.h>
#include <unistd.h>
#include "./pkg/xdxml/xdxml.h"

int main(){
    void *handle = NULL; 
    char name[XDXML_DEVICE_NAME_BUFFER_SIZE];
    xdxml_return_t result;
    unsigned int device_count = 0;

    xdxml_return_t (*xdxml_init)(void) = NULL;
    xdxml_return_t (*xdxml_device_get_count)(unsigned int*) = NULL;
    xdxml_return_t (*xdxml_device_get_handle_by_index)(unsigned int, xdx_device_t *) = NULL;
    xdxml_return_t (*xdxml_device_get_product_name)(xdx_device_t, char*) = NULL;
    xdxml_return_t (*xdxml_device_get_minor_number)(xdx_device_t, int *) = NULL;
    xdxml_return_t (*xdxml_shutdown)(void) = NULL;

    handle = dlopen("/opt/xdxgpu/lib/x86_64-linux-gnu/libxdxgpu-ml.so",RTLD_LAZY);
    if (!handle){
        printf("Could not find lib\n");
        return -1;
    }

    xdxml_init = (xdxml_return_t(*)(void))dlsym(handle,"xdxml_init");

    xdxml_device_get_count = (xdxml_return_t(*)(unsigned int*))dlsym(handle,"xdxml_device_get_count");
    if(!xdxml_device_get_count){
        printf("Open failed: There no function named xdxml_device_get_count in so.\n");
        return 0;
    }

    xdxml_device_get_handle_by_index = (xdxml_return_t(*)(unsigned int, xdx_device_t *))dlsym(handle,"xdxml_device_get_handle_by_index");
    if(!xdxml_device_get_handle_by_index){
        printf("Open failed: There no function named xdxml_device_get_handle_by_index in so.\n");
        return 0;
    }

    xdxml_device_get_minor_number = (xdxml_return_t(*)(xdx_device_t, int*))dlsym(handle,"xdxml_device_get_minor_number");
    if(!xdxml_device_get_minor_number){
        printf("Open failed: There no function named xdxml_device_get_minor_number in so.\n");
        return 0;
    }

    xdxml_device_get_product_name = (xdxml_return_t(*)(xdx_device_t, char*))dlsym(handle,"xdxml_device_get_product_name"); 
    xdxml_shutdown = (xdxml_return_t(*)(void))dlsym(handle,"xdxml_init");

    result = (*xdxml_init)();
    result = (*xdxml_device_get_count)(&device_count);

    printf("\nmain:device count is: %d\n",device_count);
    printf("Listing devices:\n");

    for (int i = 0; i < device_count; i++)
    {
        //xdx_device_t device = (xdx_device_t)malloc(sizeof(struct xdx_device_st));
        xdx_device_t device;
        result = (*xdxml_device_get_handle_by_index)(i, &device);
        if (XDXML_SUCCESS != result)
        { 
            printf("Failed to get handle.");
            return -1;
        }
        result = (*xdxml_device_get_product_name)(device, name);
        if (XDXML_SUCCESS != result)
        { 
            printf("Failed to get name of device.");
            return -1;
        }
        printf("\nproduct_name is : %s\n", name);

        int minor_number;
        result = (*xdxml_device_get_minor_number)(device, &minor_number);
        if (XDXML_SUCCESS != result)
        {
            printf("Failed to get minor_number of device.");
            return -1;
        }
        printf("\nminor_number is : %d\n", minor_number);
    }
       
    result = (*xdxml_shutdown)();
    dlclose(handle);
    return 0;
}

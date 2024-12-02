import { storage } from 'wxt/storage';

export async function setStorageItem(key:string, value:any) {
  return await storage.setItem("local:"+key,value )
}

export async function getStorageItem(key:string)  {
  return await storage.getItem("local:"+key ) as any
}

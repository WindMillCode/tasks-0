import { Tabs } from 'expo-router';
import React from 'react';
import { Platform } from 'react-native';
import { useWMLTheme } from "@/constants/Theme";


export default function WMLTemplateLayout() {
  const {colorScheme,theme} = useWMLTheme();

  return (
    <Tabs
      screenOptions={{
        headerShown: false,

        tabBarStyle: Platform.select({
          default: {
            display:"none"
          },
        }),
      }}>
      <Tabs.Screen
        name="index"
      />

    </Tabs>
  );
}

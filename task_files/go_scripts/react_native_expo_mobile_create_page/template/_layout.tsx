import { Tabs } from 'expo-router';
import React from 'react';
import { Platform } from 'react-native';
import { useColorScheme } from '@/hooks/useColorScheme';

export default function WMLTemplateLayout() {
  const colorScheme = useColorScheme();

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

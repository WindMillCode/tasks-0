import { StyleSheet } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context'


export default function WMLTemplateScreen() {

  return (
    <SafeAreaView style={styles.container}>

    </SafeAreaView>
  );
}

export const styles = StyleSheet.create({
  container: {
    flex: 1
  },
  mainView: {
    padding: 20
  },
});


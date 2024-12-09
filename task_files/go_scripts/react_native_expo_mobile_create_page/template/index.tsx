import { StyleSheet,Appearance } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context'
import { MyColors } from '@/constants/Colors';

export default function WMLTemplateScreen() {
  const colorScheme = Appearance.getColorScheme()
  const theme = colorScheme === 'dark' ? MyColors.dark : MyColors.light
  const styles = createStyles({
    colorScheme,
    theme
  });
  return (
    <SafeAreaView style={styles.container}>

    </SafeAreaView>
  );
}


function createStyles(props) {

  return StyleSheet.create({
    container: {
      flex: 1,
      backgroundColor:props.theme.background
    },
    mainView: {
      padding: 20
    },
  });
}



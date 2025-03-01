import { StyleSheet } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import {useWMLTranslation } from '@/constants/Global';
import { useWMLTheme } from '@/constants/Theme';

export default function WMLTemplatePage() {
  const {isDark,theme} = useWMLTheme();
  const styles = createStyles({
    isDark,
    theme
  });
  const { t} = useWMLTranslation();

  return (
    <SafeAreaView style={styles.container}>

    </SafeAreaView>
  );
}


function createStyles(props) {
  let {isDark,theme} = props
  return StyleSheet.create({
    label:{
      color:theme.text
    },
    container: {
      flex: 1,
      backgroundColor:theme.background
    },
    mainView: {
      padding: 20
    },
  });
}


